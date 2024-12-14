package unit

import "sync"

type SyncStreamSet[UniqueIndex, Index comparable, V any] struct {
	sync.Mutex
	*StreamSet[UniqueIndex, Index, V]
}

func NewSyncStreamSet[UniqueIndex, Index comparable, V any](reserveCount int) *SyncStreamSet[UniqueIndex, Index, V] {
	return &SyncStreamSet[UniqueIndex, Index, V]{StreamSet: NewStreamSet[UniqueIndex, Index, V](reserveCount)}
}

func (r *SyncStreamSet[UniqueIndex, Index, V]) Get(hash UniqueIndex) (v V, ok bool) {
	r.Lock()
	defer r.Unlock()
	return r.StreamSet.Get(hash)
}

func (r *SyncStreamSet[UniqueIndex, Index, V]) Push(hash UniqueIndex, height Index, v V) bool {
	r.Lock()
	defer r.Unlock()
	return r.StreamSet.Push(hash, height, v)
}
