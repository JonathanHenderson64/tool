package unit

import "sync"

type SyncOrderMap[K comparable, V any] struct {
	sync.Mutex
	*OrderMap[K, V]
}

func NewSyncOrderMap[K comparable, V any]() *SyncOrderMap[K, V] {
	return &SyncOrderMap[K, V]{OrderMap: NewOrderMap[K, V]()}
}

func (r *SyncOrderMap[K, V]) Put(k K, v V) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.OrderMap.Put(k, v)
}

func (r *SyncOrderMap[K, V]) Del(key K) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.OrderMap.Del(key)
}

func (r *SyncOrderMap[K, V]) Get(key K) (v V, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.OrderMap.Get(key)
}
