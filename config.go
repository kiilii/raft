package raft

type Config struct {
	// LHost localhost ip:port
	LHost string

	// peers remote host array
	Peers []string
}
