package raft

type Storage interface {
	Set(key string, value []byte)

	Get(key string) ([]byte, bool)

	Del(key string) ([]byte, bool)

	Has(key string) bool

	IsEmpty() bool
}
