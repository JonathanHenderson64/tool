package unit

type (
	ListNode[T comparable] struct {
		data T
		prev *ListNode[T]
		next *ListNode[T]
	}

	List[T comparable] struct {
		count int
		head  *ListNode[T]
		tail  *ListNode[T]
	}
)

func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

func (r *List[T]) Count() int {
	return r.count
}

func (r *List[T]) del(node *ListNode[T]) {
	if nil != node.prev {
		node.prev.next = node.next
	} else {
		r.head = node.next
	}
	if nil != node.next {
		node.next.prev = node.prev
	} else {
		r.tail = node.prev
	}
	r.count--
}

func (r *List[T]) Del(node *ListNode[T]) {
	r.del(node)
}

func (r *List[T]) pop(fromHead bool) (val *ListNode[T], ok bool) {
	if 0 == r.count {
		return
	}
	node := r.tail
	if fromHead {
		node = r.head
	}
	r.del(node)
	return node, true
}

func (r *List[T]) PopNodeFromHead() (val *ListNode[T], ok bool) {
	return r.pop(true)
}
func (r *List[T]) PopNodeFromTail() (val *ListNode[T], ok bool) {
	return r.pop(false)
}

func (r *List[T]) PopDataFromHead() (val T, ok bool) {
	if node, exist := r.PopNodeFromHead(); exist {
		return node.data, exist
	}
	return
}
func (r *List[T]) PopDataFromTail() (val T, ok bool) {
	if node, exist := r.PopNodeFromTail(); exist {
		return node.data, exist
	}
	return
}

func (r *List[T]) GetNodeFromHead() (val *ListNode[T], ok bool) {
	if 0 == r.count {
		return
	}
	return r.head, true
}
func (r *List[T]) GetNodeFromTail() (val *ListNode[T], ok bool) {
	if 0 == r.count {
		return
	}
	return r.tail, true
}

func (r *List[T]) GetDataFromHead() (val T, ok bool) {
	if 0 == r.count {
		return
	}
	return r.head.data, true
}
func (r *List[T]) GetDataFromTail() (val T, ok bool) {
	if 0 == r.count {
		return
	}
	return r.tail.data, true
}

func (r *List[T]) push(node *ListNode[T], fromTail bool) *ListNode[T] {
	if 0 == r.count {
		r.head = node
		r.tail = node
		r.count++
		return node
	}

	if fromTail {
		node.prev = r.tail
		r.tail.next = node
		r.tail = node
		r.count++
	} else {
		node.next = r.head
		r.head.prev = node
		r.head = node
		r.count++
	}
	return node
}

func (r *List[T]) PushDataFromHead(val T) (node *ListNode[T]) {
	return r.push(&ListNode[T]{data: val}, false)
}

func (r *List[T]) PushDataFromTail(val T) (node *ListNode[T]) {
	return r.push(&ListNode[T]{data: val}, true)
}

func (r *List[T]) RotateHeadToTail() {
	if 2 > r.count {
		return
	}
	head := r.head
	r.head = head.next
	r.head.prev = nil
	r.tail.next = head
	head.next = nil
	head.prev = r.tail
	r.tail = head
}

func (r *List[T]) RotateTailToHead() {
	if 2 > r.count {
		return
	}
	tail := r.tail
	r.tail = tail.prev
	r.tail.next = nil
	r.head.prev = tail
	tail.prev = nil
	tail.next = r.head
	r.head = tail

}

func (r *List[T]) FromHeadFind(op func(T) bool) (node *ListNode[T]) {
	for node = r.head; node != nil; node = node.next {
		if op(node.data) {
			return node
		}
	}
	return
}

func (r *List[T]) FromTailFind(op func(T) bool) (node *ListNode[T]) {
	for node = r.tail; node != nil; node = node.prev {
		if op(node.data) {
			return node
		}
	}
	return
}

func (r *List[T]) insert(val T, oldNode *ListNode[T], after bool) *ListNode[T] {
	node := &ListNode[T]{data: val}
	if after {
		node.prev = oldNode
		node.next = oldNode.next
		if r.tail == oldNode {
			r.tail = node
		}
	} else {
		node.next = oldNode
		node.prev = oldNode.prev
		if r.head == oldNode {
			r.head = node
		}
	}
	if node.prev != nil {
		node.prev.next = node
	}
	if node.next != nil {
		node.next.prev = node
	}
	r.count++
	return node
}

func (r *List[T]) InsertPrev(val T, oldNode *ListNode[T]) *ListNode[T] {
	return r.insert(val, oldNode, false)
}

func (r *List[T]) InsertNext(val T, oldNode *ListNode[T]) *ListNode[T] {
	return r.insert(val, oldNode, true)
}

func (r *List[T]) pushList(o *List[T]) {
	if o.Count() <= 0 {
		return
	}

	o.head.prev = r.tail

	if nil != r.tail {
		r.tail.next = o.head
	} else {
		r.head = o.head
	}
	r.tail = o.tail
	r.count += o.Count()
}

func (r *List[T]) PushListFromTail(o *List[T]) *List[T] {
	if o == nil {
		return r
	}

	r.pushList(o)
	return r
}
func (r *List[T]) PushListFromHead(o *List[T]) *List[T] {
	if o == nil {
		return r
	}

	o.pushList(r)
	return o
}
