package raft

import (
	"context"

	pb "github.com/kiilii/raft/proto"
	"google.golang.org/grpc"
)

type GrpcTransport struct {
	s *grpc.Server
}

func NewGrpcTransport(service pb.TransportServerServer) *GrpcTransport {

	// options
	svr := grpc.NewServer()

	pb.RegisterTransportServerServer(svr, service)

	return &GrpcTransport{
		s: svr,
	}
}

type raftGrpcServer struct {
	pb.UnimplementedTransportServerServer
}

func (s *raftGrpcServer) RequestVote(ctx context.Context, in *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {

	return nil, nil
}

func (s *raftGrpcServer) AppendEntry(ctx context.Context, in *pb.AppendEntryArgs) (*pb.AppendEntryReply, error) {
	return nil, nil
}
