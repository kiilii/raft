package raft

type Config struct {
	// LHost localhost ip:port
	LHost string

	// peers remote host array
	Peers []string

	// 过期时间
	Timeout      int
	ReadTimeout  int
	WriteTimeout int
}
