package server

import "github.com/hashicorp/raft"

type raftNode struct {
	*raft.Raft
}

func newRaftNode(conf RaftConfig) (*raftNode, error) {
	panic("TODO")
}

func (r *raftNode) start() error {
	panic("TODO")
}
