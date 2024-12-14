package unit

import (
	"sync"
)

type SyncQueue[T comparable] struct {
	running bool
	max     int
	queue   *List[T]
	mutex   sync.Mutex
	empty   *sync.Cond
	full    *sync.Cond
}

func NewSyncQueue[T comparable](max int) *SyncQueue[T] {
	r := &SyncQueue[T]{
		max:     max,
		running: true,
		queue:   NewList[T](),
	}
	r.empty = sync.NewCond(&r.mutex)
	r.full = sync.NewCond(&r.mutex)
	return r
}

func (r *SyncQueue[T]) count() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.queue.Count()
}

func (r *SyncQueue[T]) isFull() bool {
	if r.max > 0 && r.queue.Count() >= r.max {
		return true
	}
	return false
}

func (r *SyncQueue[T]) isEmpty() bool {
	return r.queue.Count() == 0
}

func (r *SyncQueue[T]) pop() (val T, ok bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for r.isEmpty() && r.running {
		r.empty.Wait()
	}

	val, ok = r.queue.PopDataFromHead()
	if r.max > 0 {
		r.full.Signal()
	}
	return
}

func (r *SyncQueue[T]) insert(val T, priority bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for r.isFull() && r.running {
		r.full.Wait()
	}

	if r.running {
		if priority {
			r.queue.PushDataFromHead(val)
		} else {
			r.queue.PushDataFromTail(val)
		}
		r.empty.Signal()
	}
}

func (r *SyncQueue[T]) add(val T) {
	r.insert(val, true)
}

func (r *SyncQueue[T]) push(val T) {
	r.insert(val, false)
}

func (r *SyncQueue[T]) insertMore(l *List[T], priority bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for r.isFull() && r.running {
		r.full.Wait()
	}

	if r.running {
		if priority {
			r.queue = r.queue.PushListFromHead(l)
		} else {
			r.queue = r.queue.PushListFromTail(l)
		}
		r.empty.Broadcast()
	}
}

func (r *SyncQueue[T]) addMore(l *List[T]) {
	r.insertMore(l, true)
}

func (r *SyncQueue[T]) pushMore(l *List[T]) {
	r.insertMore(l, false)
}

func (r *SyncQueue[T]) stop() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.running = false
	r.empty.Broadcast()
	if r.max > 0 {
		r.full.Broadcast()
	}
}
