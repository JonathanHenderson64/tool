package unit

type MultiMap[K1, K2 comparable, V any] struct {
	m map[K1]map[K2]V
}

func NewMultiMap[K1, K2 comparable, V any]() *MultiMap[K1, K2, V] {
	return &MultiMap[K1, K2, V]{m: make(map[K1]map[K2]V)}
}

func (r *MultiMap[K1, K2, V]) Push(k1 K1, k2 K2, v V) {
	if m, ok := r.m[k1]; ok {
		m[k2] = v
		r.m[k1] = m
	} else {
		r.m[k1] = map[K2]V{k2: v}
	}
}

func (r *MultiMap[K1, K2, V]) Get(k1 K1, k2 K2) (v V, ok bool) {
	if m, exist := r.m[k1]; exist {
		v, ok = m[k2]
	}
	return
}

func (r *MultiMap[K1, K2, V]) Del(k1 K1, k2 K2) (v V, ok bool) {
	m, exist := r.m[k1]
	if !exist {
		return
	}

	v, ok = m[k2]
	if !ok {
		return
	}

	delete(m, k2)
	if len(m) == 0 {
		delete(r.m, k1)
	} else {
		r.m[k1] = m
	}
	return
}
