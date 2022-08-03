package raft

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	c        *Config
	shutdown bool

	mu    sync.RWMutex
	conns map[string]*net.TCPConn
}

func New() *Server {
	s := &Server{shutdown: true, c: &Config{}}

	s.connenctPeers()
	go s.listen()

	return s
}

// connenctPeers 连接对等节点
func (s *Server) connenctPeers() {
	for _, p := range s.c.Peers {
		raddr, err := net.ResolveTCPAddr("tcp", p)
		if err != nil {
			panic(err) // 配置错误直接崩溃
		}

		conn, err := net.DialTCP("tcp", nil, raddr)
		if err != nil {
			fmt.Println("") // 连接失败打印信息
		}

		s.mu.Lock()
		s.conns[conn.RemoteAddr().String()] = conn
		s.mu.Unlock()
	}
}

// listen 开启服务监听
func (s *Server) listen() {
	tp := NewTransport(s.c)

	// 开启服务监听
	for s.shutdown {
		conn, err := tp.Accept()
		if err != nil {
			fmt.Printf("exception connnet! error(%v)", err)
		}

		if c, has := s.conns[conn.RemoteAddr().String()]; has {
			fmt.Printf("exception connnet! error(%v)", err)
			if err := c.Close(); err != nil {
				fmt.Printf("close connent error! %v", err)
			}
		}
		s.mu.Lock()
		s.conns[conn.RemoteAddr().String()] = conn
		s.mu.Unlock()
	}
}
