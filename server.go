package raft

import (
	"context"

	pb "github.com/kiilii/raft/proto"
)

// Server raft 内部服务。用于与其他 peers 通讯
type Server struct {
	c *Config

	// commitLogs chan<-

	pb.UnimplementedTransportServerServer
}

func NewServer(c *Config) *Server {
	return &Server{c: c}
}

func (s *Server) RequestVote(context.Context, *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
	panic("sad")
}
func (s *Server) AppendEntry(context.Context, *pb.AppendEntryArgs) (*pb.AppendEntryReply, error) {
	panic("mad")
}
