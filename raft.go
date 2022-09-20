package raft

type RaftServer struct {
	timeElapsed               int64
	requestTime               int64
	electionTimeout           int64
	voting_cfg_change_log_idx int64
	// current_term = 0;
	// voted_for = -1;
	// timeout_elapsed = 0;
	// request_timeout = 200;
	// election_timeout = 1000;

	wal *WAL

	// 表示当服务状态
	state RaftState

	currentTerm int64

	votedFor int64

	// 当前集群 leader 节点
	currentLeader *Node
	// 自身节点
	node *Node
	// 集群中的所有节点
	nodes map[uint64]*Node

	// log commited id
	commitIdx uint64
}

func New(c *Config) *RaftServer {
	var x = &RaftServer{
		currentTerm:               0,
		votedFor:                  -1,
		timeElapsed:               c.Timeout,
		requestTime:               c.RequestTimeout,
		electionTimeout:           c.ElectionTimeout,
		wal:                       NewWAL(),
		voting_cfg_change_log_idx: -1,
	}

	return x
}

func (rs *RaftServer) SetState(state RaftState) {
	if state == RaftStateLeader {
		rs.currentLeader = rs.node
	}
	rs.state = state
}
func (rs *RaftServer) GetState() RaftState { return rs.state }

func (rs *RaftServer) GetNode(nodeid uint64) *Node {
	for _, node := range rs.nodes {
		if node.GetNodeID() == nodeid {
			return node
		}
	}
	return nil
}

// GetLocalNode 获取本地节点
func (rs *RaftServer) GetLocalNode() *Node {
	for _, node := range rs.nodes {
		if rs.node.GetNodeID() == node.GetNodeID() {
			return node
		}
	}
	return nil
}

func (rs *RaftServer) GetNodeByIdx(idx uint64) *Node {
	for i, node := range rs.nodes {
		if i == idx {
			return node
		}
	}
	return nil
}

func (rs *RaftServer) GetCommitIdx() uint64 { return rs.commitIdx }

func (rs *RaftServer) SetCommitIdx(idx uint64) {
	// TODO:: 缺少实际日志记录校验:
	// 记录必须提交，持久化后的记录才可更新 commitIdx
	// if rs.log.commitedId >= idx
	if rs.commitIdx <= idx {
		rs.commitIdx = idx
	}
	rs.commitIdx = idx
}

func (rs *RaftServer) GetCurrentLeader() (uint64, bool) {
	if rs.currentLeader != nil {
		return rs.currentLeader.GetNodeID(), true
	}
	return 0, false
}

func (rs *RaftServer) GetCurrentLeaderNode() *Node { return rs.currentLeader }

func (rs *RaftServer) IsFollower() bool { return rs.state == RaftStateFollower }

func (rs *RaftServer) IsCandidate() bool { return rs.state == RaftStateCandidate }

func (rs *RaftServer) IsLeader() bool { return rs.state == RaftStateLeader }

// SetCurrentTerm 设置当前 leader 任期
func (rs *RaftServer) SetCurrentTerm(term int64) (err error) {
	if rs.currentTerm <= term {
		// TODO::
		rs.currentTerm = term
	}
	return err
}

func (rs *RaftServer) GetCurrentTerm() int64 { return rs.currentTerm }
