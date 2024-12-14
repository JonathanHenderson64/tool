package unit

import "sync"

type SyncMapArr[K comparable, V any] struct {
	sync.Mutex
	*MapArr[K, V]
}

func NewSyncMapArr[K comparable, V any]() *SyncMapArr[K, V] {
	return &SyncMapArr[K, V]{MapArr: NewMapArr[K, V]()}
}

func (r *SyncMapArr[K, V]) Push(k K, v V) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.MapArr.Push(k, v)
}

func (r *SyncMapArr[K, V]) Get(k K) ([]V, bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.MapArr.Get(k)
}

func (r *SyncMapArr[K, V]) Del(k K) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.MapArr.Del(k)
}

func (r *SyncMapArr[K, V]) Count() int {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.MapArr.Count()
}
