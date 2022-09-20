package raft

const (
	// node state
	_ = iota << 0
	RaftNodeVotedForMe
	RaftNodeVoting
	RaftNodeVotingCommited
	RaftNodeInactive
)
