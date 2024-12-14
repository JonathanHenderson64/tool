package unit

import "sync"

type SyncList[T comparable] struct {
	sync.Mutex
	*List[T]
}

func NewSyncList[T comparable]() *SyncList[T] {
	return &SyncList[T]{List: NewList[T]()}
}

func (r *SyncList[T]) Count() int {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.count
}

func (r *SyncList[T]) Del(node *ListNode[T]) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.List.Del(node)
}
func (r *SyncList[T]) PopNodeFromHead() (val *ListNode[T], ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PopNodeFromHead()
}
func (r *SyncList[T]) PopNodeFromTail() (val *ListNode[T], ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PopNodeFromTail()
}

func (r *SyncList[T]) PopDataFromHead() (val T, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PopDataFromHead()
}
func (r *SyncList[T]) PopDataFromTail() (val T, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PopDataFromTail()
}

func (r *SyncList[T]) GetNodeFromHead() (val *ListNode[T], ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.GetNodeFromHead()
}
func (r *SyncList[T]) GetNodeFromTail() (val *ListNode[T], ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.GetNodeFromTail()
}

func (r *SyncList[T]) GetDataFromHead() (val T, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.GetDataFromHead()
}
func (r *SyncList[T]) GetDataFromTail() (val T, ok bool) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.GetDataFromTail()
}

func (r *SyncList[T]) PushDataFromHead(val T) (node *ListNode[T]) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PushDataFromHead(val)
}

func (r *SyncList[T]) PushDataFromTail(val T) (node *ListNode[T]) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PushDataFromTail(val)
}

func (r *SyncList[T]) RotateHeadToTail() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.List.RotateHeadToTail()
}

func (r *SyncList[T]) RotateTailToHead() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.List.RotateTailToHead()
}

func (r *SyncList[T]) FromHeadFind(op func(T) bool) (node *ListNode[T]) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.FromHeadFind(op)
}

func (r *SyncList[T]) FromTailFind(op func(T) bool) (node *ListNode[T]) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.FromTailFind(op)
}

func (r *SyncList[T]) InsertPrev(val T, oldNode *ListNode[T]) *ListNode[T] {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.InsertPrev(val, oldNode)
}

func (r *SyncList[T]) InsertNext(val T, oldNode *ListNode[T]) *ListNode[T] {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.InsertNext(val, oldNode)
}

func (r *SyncList[T]) PushList(o *List[T]) *List[T] {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PushListFromTail(o)
}
func (r *SyncList[T]) AddList(o *List[T]) *List[T] {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.List.PushListFromHead(o)
}
