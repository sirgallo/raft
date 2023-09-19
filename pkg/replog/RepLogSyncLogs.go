package replog

import "github.com/sirgallo/raft/pkg/system"


//=========================================== Sync Logs


/*
	Sync Logs:
		helper method for handling syncing followers who have inconsistent logs

		while unsuccessful response:
			send AppendEntryRPC to follower with logs starting at the follower's NextIndex
			if error: return false, error
			on success: return true, nil --> the log is now up to date with the leader
*/

func (rlService *ReplicatedLogService) SyncLogs(host string) (bool, error) {
	s, _ := rlService.Systems.Load(host)
	sys := s.(*system.System)

	sys.SetStatus(system.Busy)

	conn, connErr := rlService.ConnectionPool.GetConnection(sys.Host, rlService.Port)
	if connErr != nil {
		rlService.Log.Error("Failed to connect to", sys.Host + rlService.Port, ":", connErr.Error())
		return false, connErr
	}

	for {
		earliestLog, earliestErr := rlService.CurrentSystem.WAL.GetEarliest()
		if earliestErr != nil { return false, earliestErr }

		preparedEntries, prepareErr := rlService.PrepareAppendEntryRPC(sys.NextIndex, false)
		if prepareErr != nil { return false, prepareErr }

		req := ReplicatedLogRequest{
			Host: sys.Host,
			AppendEntry: preparedEntries,
		}

		res, rpcErr := rlService.clientAppendEntryRPC(conn, sys, req)
		if rpcErr != nil { return false, rpcErr }
		
		if res.Success {
			sys.SetStatus(system.Ready)
			rlService.ConnectionPool.PutConnection(sys.Host, conn)

			return true, nil
		} else if earliestLog != nil && sys.NextIndex < earliestLog.Index {
			rlService.SendSnapshotToSystemSignal <- sys.Host
			return true, nil
		}
	}
}