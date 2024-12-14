package unit

type StreamSet[UniqueIndex, Index comparable, V any] struct {
	hashSet      map[UniqueIndex]V           //map[hash]share
	heightMapArr *MapArr[Index, UniqueIndex] //map[height][]hash
	heightQueue  *List[Index]                //[]height
	reserveCount int
}

func NewStreamSet[UniqueIndex, Index comparable, V any](reserveCount int) *StreamSet[UniqueIndex, Index, V] {
	return &StreamSet[UniqueIndex, Index, V]{
		hashSet:      make(map[UniqueIndex]V),
		heightMapArr: NewMapArr[Index, UniqueIndex](),
		heightQueue:  NewList[Index](),
		reserveCount: reserveCount,
	}
}

func (r *StreamSet[UniqueIndex, Index, V]) Get(hash UniqueIndex) (v V, ok bool) {
	v, ok = r.hashSet[hash]
	return
}

func (r *StreamSet[UniqueIndex, Index, V]) Push(hash UniqueIndex, height Index, v V) bool {
	if _, ok := r.hashSet[hash]; ok {
		return true
	}

	r.hashSet[hash] = v
	if _, ok := r.heightMapArr.Get(height); !ok {
		r.heightQueue.PushDataFromTail(height)
	}
	r.heightMapArr.Push(height, hash)

	for r.heightQueue.Count() > r.reserveCount {
		delHeight, ok := r.heightQueue.PopDataFromHead()
		if !ok {
			return false
		}

		delHashes, exist := r.heightMapArr.Get(delHeight)
		if !exist {
			return false
		}

		for _, delHash := range delHashes {
			delete(r.hashSet, delHash)
		}
		r.heightMapArr.Del(delHeight)
	}
	return false
}
