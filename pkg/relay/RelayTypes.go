package relay

import "sync"
import "time"

import "github.com/sirgallo/raft/pkg/relayrpc"
import "github.com/sirgallo/raft/pkg/logger"
import "github.com/sirgallo/raft/pkg/connpool"
import "github.com/sirgallo/raft/pkg/statemachine"
import "github.com/sirgallo/raft/pkg/system"


type RelayOpts struct {
	Port int
	ConnectionPool *connpool.ConnectionPool

	CurrentSystem *system.System
	SystemsList []*system.System
	Systems *sync.Map
}

type RelayService struct {
	relayrpc.UnimplementedRelayServiceServer
	Port string
	ConnectionPool *connpool.ConnectionPool

	// Persistent State
	CurrentSystem *system.System
	Systems *sync.Map

	// Module Level State
	RelayChannel chan statemachine.StateMachineOperation
	RelayedAppendLogSignal chan statemachine.StateMachineOperation

	Log clog.CustomLog
}


const NAME = "Relay"
const RPCTimeout = 50 * time.Millisecond
const RelayChannelBuffSize = 100000
const FailedBuffSize = 10000