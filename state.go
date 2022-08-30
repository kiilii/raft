package raft

import (
	"sync/atomic"
)

type stateType int32

const (
	Follower stateType = iota
	Candidate
	Leader
)

type State struct {
	state stateType
}

func (s *State) GetState(st stateType) stateType {
	return stateType(atomic.LoadInt32((*int32)(&s.state)))
}

func (s *State) SetState(st stateType) {
	atomic.StoreInt32((*int32)(&s.state), int32(st))
}

func (s *State) IsFollower() bool  { return s.state == Follower }
func (s *State) IsCandidate() bool { return s.state == Candidate }
func (s *State) IsLeader() bool    { return s.state == Leader }
