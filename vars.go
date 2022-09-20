package raft

const (
	// node state
	_ = iota << 0
	RaftNodeVotedForMe
	RaftNodeVoting
	RaftNodeVotingCommited
	RaftNodeInactive
)

type RaftState int

const (
	RAFT_STATE_NONE RaftState = iota + 1
	RaftStateFollower
	RaftStateCandidate
	RaftStateLeader
)
