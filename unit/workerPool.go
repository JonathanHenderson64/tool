package unit

import (
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type WorkerPool struct {
	latch     sync.WaitGroup
	panicked  chan bool
	running   atomic.Bool
	onceClose sync.Once
	queue     *SyncQueue[AnyEvent]
}

type AnyEvent interface {
	Callback()
}

type AnyLogger interface {
	Panic(v ...any)
}

var g_logger AnyLogger

func (r *WorkerPool) SetLogger(logger AnyLogger) {
	g_logger = logger
}
func (r *WorkerPool) run() {
	r.latch.Add(1)
	defer func() {
		if err := recover(); nil != err {
			if g_logger != nil {
				g_logger.Panic(string(debug.Stack()))
			}
			r.panicked <- true
		}
		r.latch.Done()
	}()

	for {
		if task, ok := r.take(); ok {
			task.Callback()
		} else if !r.running.Load() {
			return
		}
	}
}
func NewWorkerPool(queueSize, threads int) *WorkerPool {
	r := &WorkerPool{queue: NewSyncQueue[AnyEvent](queueSize), panicked: make(chan bool)}
	r.running.Store(true)
	for i := 0; i < threads; i++ {
		go r.run()
	}
	go func() {
		for r.running.Load() {
			select {
			case <-r.panicked:
				if r.running.Load() {
					go r.run()
				}
			}
		}
	}()
	return r
}

func (r *WorkerPool) Close() {
	r.onceClose.Do(func() {
		r.running.Store(false)
		close(r.panicked)
		r.queue.stop()
		r.latch.Wait()
	})
}

func (r *WorkerPool) Add(task AnyEvent) {
	if r.running.Load() {
		r.queue.add(task)
	}
}

func (r *WorkerPool) Push(task AnyEvent) {
	if r.running.Load() {
		r.queue.push(task)
	}

}

func (r *WorkerPool) take() (AnyEvent, bool) {
	return r.queue.pop()
}
