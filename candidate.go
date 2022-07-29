package raft

import (
	"fmt"
	"runtime"
)

type Candidate struct {
	Status bool // on off
}

func NewCandidate() *Candidate {
	return &Candidate{
		Status: true,
	}
}

func (c *Candidate) Watch() error {
	var err error

	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
	}()

	for c.Status {
		// c.svr.RequestVote()
	}

	return err
}

func (c *Candidate) Stop() {
	c.Status = false
}

func (c *Candidate) Start() {
	c.Status = true
}
