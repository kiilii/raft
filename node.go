package raft

// 保证 node 之间互通

type Node struct {
	NextIdx  uint64
	MatchIdx uint64

	// node state
	flag int
	id   uint64
}

func NewNode(id uint64) *Node {
	return &Node{
		id: id,
	}
}

func (n *Node) GetNodeID() uint64 {
	return n.id
}

func (n *Node) GetNextIdx() uint64 {
	return n.NextIdx
}
func (n *Node) SetNextIdx(idx uint64) {
	if idx < 1 {
		idx = 1
	}
	n.NextIdx = idx
}

// VotedForMe ...
func (n *Node) SetVotedForMe(isVote bool) {
	if isVote {
		n.flag |= RaftNodeVotedForMe
	} else {
		n.flag &= ^RaftNodeVotedForMe
	}
}
func (n *Node) SetVoting(isVoting bool) {
	if isVoting {
		n.flag |= RaftNodeVoting
	} else {
		n.flag &= ^RaftNodeVoting
	}
}
func (n *Node) SetActive(isAvtive bool) {
	if !isAvtive {
		n.flag |= RaftNodeInactive
	} else {
		n.flag &= ^RaftNodeInactive
	}
}
func (n *Node) SetVotingCommited(isVotingCommited bool) {
	if isVotingCommited {
		n.flag |= RaftNodeVotingCommited
	} else {
		n.flag &= ^RaftNodeVotingCommited
	}
}

func (n *Node) IsVoting() bool         { return (n.flag & RaftNodeVoting) != 0 }
func (n *Node) IsVotedForMe() bool     { return (n.flag & RaftNodeVotedForMe) != 0 }
func (n *Node) IsActive() bool         { return (n.flag & RaftNodeInactive) == 0 }
func (n *Node) IsVotingCommited() bool { return (n.flag & RaftNodeVotingCommited) != 0 }
