package forwardresp

import "context"

import "github.com/sirgallo/raft/pkg/forwardresprpc"
import "github.com/sirgallo/raft/pkg/statemachine"
import "github.com/sirgallo/raft/pkg/system"
import "github.com/sirgallo/raft/pkg/utils"


//=========================================== Forward Response Client


/*
	Forward Resp Client RPC:
		used by leader to relay responses back to the follower that relayed the request
		1.) encode the response to string
		2.) attempt forward using exponential backoff
		3.) if response is not processed, return error
		4.) otherwise, success
*/

func (frService *ForwardRespService) ForwardRespClientRPC(resp *statemachine.StateMachineResponse) (*forwardresprpc.ForwardRespResponse, error){
	strResp, encErr := utils.EncodeStructToString[*statemachine.StateMachineResponse](resp)
	if encErr != nil { 
		frService.Log.Error("error encoding log struct to string") 
		return nil, encErr
	}
	
	relayReq := &forwardresprpc.ForwardRespRequest{
		Host: frService.CurrentSystem.Host,
		StateMachineResponse: strResp,
	}

	conn, connErr := frService.ConnectionPool.GetConnection(resp.RequestOrigin, frService.Port)
	if connErr != nil { 
		frService.Log.Error("Failed to connect to", resp.RequestOrigin + frService.Port, ":", connErr.Error()) 
		return nil, connErr
	}

	client := forwardresprpc.NewForwardRespServiceClient(conn)

	forwardRespRPC := func() (*forwardresprpc.ForwardRespResponse, error) {
		ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
		defer cancel()
		
		res, err := client.ForwardRespRPC(ctx, relayReq)
		if err != nil { return utils.GetZero[*forwardresprpc.ForwardRespResponse](), err }
		if ! res.ProcessedRequest { 
			frService.Log.Error("forward request not processed")
			return utils.GetZero[*forwardresprpc.ForwardRespResponse](), err
		}

		return res, nil
	}
	
	maxRetries := 5
	expOpts := utils.ExpBackoffOpts{ MaxRetries: &maxRetries, TimeoutInMilliseconds: 1 }
	expBackoff := utils.NewExponentialBackoffStrat[*forwardresprpc.ForwardRespResponse](expOpts)
			
	res, err := expBackoff.PerformBackoff(forwardRespRPC)
	if err != nil { 
		frService.Log.Warn("system", resp.RequestOrigin, "unreachable, setting status to dead")

		sys, ok := frService.Systems.Load(resp.RequestOrigin)
		if ok { sys.(*system.System).SetStatus(system.Dead) }

		frService.ConnectionPool.CloseConnections(resp.RequestOrigin)
		
		return nil, err
	}

	return res, nil
}