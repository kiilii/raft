package raft

import (
	"log"
	"net"

	"github.com/kiilii/raft/proto"
	"google.golang.org/grpc"
)

type RaftServer struct {
	c *Config

	ServerID string
	PeersID  []string

	quit chan<- interface{}

	// 服务监听本体
	server *grpc.Server
	listen net.Listener
}

func New(c *Config) *RaftServer {
	var lis, err = net.Listen("tcp", c.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rs := &RaftServer{
		c:      c,
		quit:   make(chan<- interface{}),
		listen: lis,
	}

	// register grpc
	nodeServer := grpc.NewServer()
	proto.RegisterPeerServer(nodeServer, NewServer(rs.c))
	rs.server = nodeServer

	if rs.server.Serve(lis); err != nil {
		panic(err)
	}

	// 进入选举流程
	return rs
}
