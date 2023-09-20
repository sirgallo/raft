package service

import "sync"

import "github.com/sirgallo/raft/pkg/connpool"
import "github.com/sirgallo/raft/pkg/forwardresp"
import "github.com/sirgallo/raft/pkg/httpservice"
import "github.com/sirgallo/raft/pkg/leaderelection"
import "github.com/sirgallo/raft/pkg/replog"
import "github.com/sirgallo/raft/pkg/relay"
import "github.com/sirgallo/raft/pkg/snapshot"
import "github.com/sirgallo/raft/pkg/statemachine"
import "github.com/sirgallo/raft/pkg/system"


type RaftPortOpts struct {
	HTTPService int
	LeaderElection int
	ReplicatedLog int
	Relay int
	ForwardResp int
	Snapshot int
}

type RaftServiceOpts struct {
	Protocol string
	Ports RaftPortOpts
	SystemsList []*system.System
	ConnPoolOpts connpool.ConnectionPoolOpts
}

type RaftService struct {
	Protocol string
	Ports RaftPortOpts
	CurrentSystem *system.System
	Systems *sync.Map

	HTTPService *httpservice.HTTPService
	LeaderElection *leaderelection.LeaderElectionService
	ReplicatedLog *replog.ReplicatedLogService
	Relay *relay.RelayService
	ForwardResp *forwardresp.ForwardRespService
	Snapshot *snapshot.SnapshotService

	CommandChannel chan statemachine.StateMachineOperation
}


const DefaultCommitIndex = -1
const DefaultLastApplied = -1
const CommandChannelBuffSize = 100000