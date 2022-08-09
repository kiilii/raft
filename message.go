package raft

import (
	"bytes"
	"encoding/binary"
)

const (
	_magicNumber = 0x77777777
)

const (
	CandidateCommand = iota + 1
	OtherCommand
)

type Message struct {
	header [16]byte
	// command Raft 协议通讯指令
	command uint16
	// body 消息体
	body []byte // protobuf 内容
}

// NewMessage ...
func NewMessage(cmd uint16, body []byte) *Message {
	var m = &Message{
		command: cmd,
		body:    body,
	}
	m.initHeader()
	m.setCommand(cmd)
	m.setLength(uint32(len(body)))
	return m
}

func (m *Message) initHeader()             { binary.BigEndian.PutUint32(m.header[:4], _magicNumber) }
func (m *Message) setCommand(cmd uint16)   { binary.BigEndian.PutUint16(m.header[4:6], cmd) }
func (m *Message) setLength(length uint32) { binary.BigEndian.PutUint32(m.header[6:10], length) }

// Decode ...
func (m *Message) Decode() []byte {
	buf := bytes.NewBuffer(make([]byte, 0))

	buf.Write(m.header[:16])
	buf.Write(m.body)

	return buf.Bytes()
}

// Encode ...
func (m *Message) Encode(buf []byte) {
	m.setCommand(binary.BigEndian.Uint16(buf[4:6]))
	m.setLength(binary.BigEndian.Uint32(buf[6:10]))
	m.body = buf[16:]
}
