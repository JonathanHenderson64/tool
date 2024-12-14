package unit

type orderMapNode[K comparable, V any] struct {
	key   K
	value V
	index *ListNode[*orderMapNode[K, V]]
}

type OrderMap[K comparable, V any] struct {
	m map[K]*orderMapNode[K, V]
	l *List[*orderMapNode[K, V]]
}

func NewOrderMap[K comparable, V any]() *OrderMap[K, V] {
	return &OrderMap[K, V]{m: make(map[K]*orderMapNode[K, V]), l: NewList[*orderMapNode[K, V]]()}
}

func (r *OrderMap[K, V]) Put(k K, v V) {
	if node, ok := r.m[k]; ok {
		node.value = v
		return
	}

	node := &orderMapNode[K, V]{
		key:   k,
		value: v,
	}
	node.index = r.l.PushDataFromTail(node)
	r.m[k] = node
}

func (r *OrderMap[K, V]) Del(key K) {
	if entry, ok := r.m[key]; ok {
		r.l.del(entry.index)
		delete(r.m, key)
	}
}

func (r *OrderMap[K, V]) Get(key K) (v V, ok bool) {
	if entry, exist := r.m[key]; exist {
		return entry.value, true
	}
	return
}
