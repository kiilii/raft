package raft

import "testing"

func TestNode_SetVotedForMe(t *testing.T) {
	type fields struct {
		NextIdx  uint64
		MatchIdx uint64
		flag     int
	}
	type args struct {
		isVote bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{},
			args: args{
				isVote: true,
			},
		},
		{
			name:   "2",
			fields: fields{},
			args: args{
				isVote: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				NextIdx:  tt.fields.NextIdx,
				MatchIdx: tt.fields.MatchIdx,
				flag:     tt.fields.flag,
			}
			n.SetVotedForMe(tt.args.isVote)
			n.SetVotedForMe(!tt.args.isVote)
		})
	}
}
