package unit

import "sync"

type SyncMultiMap[K1, K2 comparable, V any] struct {
	sync.Mutex
	*MultiMap[K1, K2, V]
}

func NewSyncMultiMap[K1, K2 comparable, V any]() *SyncMultiMap[K1, K2, V] {
	return &SyncMultiMap[K1, K2, V]{MultiMap: NewMultiMap[K1, K2, V]()}
}

func (r *SyncMultiMap[K1, K2, V]) Set(k1 K1, k2 K2, v V) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.MultiMap.Set(k1, k2, v)
}

func (r *SyncMultiMap[K1, K2, V]) Get(k1 K1, k2 K2) (v V, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.MultiMap.Get(k1, k2)
}

func (r *SyncMultiMap[K1, K2, V]) Del(k1 K1, k2 K2) (v V, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.MultiMap.Del(k1, k2)
}
