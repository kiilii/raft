package raft

import (
	"log"
	"net"

	"github.com/kiilii/raft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RaftServer struct {
	c *Config

	ServerID string

	quit chan<- interface{}

	// 服务监听本体
	server *grpc.Server
	listen net.Listener
	peers  map[int]grpc.ClientConnInterface

	logger *log.Logger
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
		logger: log.Default(),
	}

	// register grpc
	nodeServer := grpc.NewServer()
	proto.RegisterPeerServer(nodeServer, NewServer(rs.c))
	rs.server = nodeServer

	// 开启服务
	if err := rs.server.Serve(lis); err != nil {
		panic(err)
	}
	// 连接各个 peers 节点
	if err := rs.ConnectAllPeers(); err != nil {
		panic(err)
	}

	return rs
}

func (rs *RaftServer) ConnectAllPeers() error {
	for _, peer := range rs.c.Peers {
		conn, err := grpc.Dial(peer, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			rs.logger.Print(err)
			continue
		}

		peerid, err := IP2Number(peer)
		if err != nil {
			rs.logger.Print(err)
			continue
		}

		if _, has := rs.peers[peerid]; !has {
			rs.peers[peerid] = conn
		}
	}
	return nil
}
