package raft

import (
	"fmt"
	"testing"

	pb "github.com/kiilii/raft/proto"
	"google.golang.org/protobuf/proto"
)

func TestNewMessage(t *testing.T) {
	var args = &pb.AppendEntryArgs{
		Term:         1,
		LeaderId:     2,
		PrevLogIndex: 3,
		PrevLogTerm:  4,
		LeaderCommit: 5,
	}

	// fmt.Println(args.Nar)

	byt, err := proto.Marshal(args)
	fmt.Println(byt, err, proto.Size(args))

	// NewMessage(CandidateCommand, )

}
