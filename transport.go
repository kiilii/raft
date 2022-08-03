package raft

import (
	"fmt"
	"net"
)

type Transport struct {
	conf     *Config
	listener *net.TCPListener
}

func NewTransport(c *Config) *Transport {
	addr, err := net.ResolveTCPAddr("tcp", c.LHost)
	if err != nil {
		panic(fmt.Sprintf("resolve addr error! error(%s)", err))
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("tcp listen failed! error(%s)", err))
	}

	return &Transport{
		conf:     c,
		listener: l,
	}
}

func (t *Transport) Accept() (*net.TCPConn, error) {
	return t.listener.AcceptTCP()
}

func (t *Transport) Stop() error {
	return t.listener.Close()
}
