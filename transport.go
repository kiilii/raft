package raft

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// transport

type Transport struct {
	conf *Config

	mu    *sync.RWMutex
	conns map[string]net.Conn

	shutdown bool

	logger *log.Logger
}

func NewTransport(c *Config) *Transport {
	tp := &Transport{
		conf:   c,
		mu:     &sync.RWMutex{},
		logger: log.Default(),
	}

	go tp.Listener()

	return tp
}

func (tp *Transport) Listener() {
	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(fmt.Sprintf("can not listener!,%+v", err))
	}
	for !tp.shutdown {
		conn, err := l.Accept()
		if err != nil {
			tp.logger.Printf("exception of accept conn. %+v", err)
			continue
		}
		tp.saveConn(conn)
	}

	tp.mu.Lock()
	defer tp.mu.Unlock()

	for _, conn := range tp.conns {
		if err := conn.Close(); err != nil {
			tp.logger.Printf("close connenct error! err(%v)", err.Error())
		}
	}

	if err := l.Close(); err != nil {
		tp.logger.Printf("close connenct error! err(%v)", err.Error())
	}
}

func (tp *Transport) Close() {
	tp.shutdown = false
}

func (tp *Transport) saveConn(conn net.Conn) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	var addr = conn.RemoteAddr().String()

	_, has := tp.conns[addr]
	if !has {
		tp.conns[addr] = conn
	} else {
		tp.logger.Printf("duplicate conn! remote(%s)", addr)
	}
}
