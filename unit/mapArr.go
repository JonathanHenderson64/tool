package unit

type MapArr[K comparable, V any] struct {
	m map[K][]V
}

func NewMapArr[K comparable, V any]() *MapArr[K, V] {
	return &MapArr[K, V]{m: make(map[K][]V)}
}

func (r *MapArr[K, V]) Push(k K, v V) {
	if sli, ok := r.m[k]; ok {
		r.m[k] = append(sli, v)
	} else {
		r.m[k] = []V{v}
	}
}

func (r *MapArr[K, V]) Get(k K) ([]V, bool) {
	v, ok := r.m[k]
	return v, ok
}

func (r *MapArr[K, V]) Del(k K) {
	delete(r.m, k)
}

func (r *MapArr[K, V]) Count() int {
	return len(r.m)
}
