package raft

type Config struct {
	// Host localhost ip:port
	Host string

	// peers remote host array
	Peers []string

	// 过期时间
	Timeout      int
	ReadTimeout  int
	WriteTimeout int
}
