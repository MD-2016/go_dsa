package hashmap

type HashMap[K, T comparable] struct {
	m map[K]T
}

func New[K, T comparable]() *HashMap[K, T] {
	return &HashMap[K, T]{m: make(map[K]T)}
}

func (hm *HashMap[K, T]) Put(key K, val T) {
	hm.m[key] = val
}

func (hm *HashMap[K, T]) Get(key K) (val T, keyFound bool) {
	val, keyFound = hm.m[key]
	return
}

func (hm *HashMap[K, T]) Remove(key K) {
	delete(hm.m, key)
}

func (hm *HashMap[K, T]) IsEmpty() bool {
	return hm.Size() == 0
}

func (hm *HashMap[K, T]) Size() int {
	return len(hm.m)
}

func (hm *HashMap[K, T]) Keys() []K {
	keys := make([]K, hm.Size())
	keyCount := 0
	for key := range hm.m {
		keys[keyCount] = key
		keyCount++
	}
	return keys
}

func (hm *HashMap[K, T]) Values() []T {
	vals := make([]T, hm.Size())
	valCount := 0
	for _, val := range hm.m {
		vals[valCount] = val
		valCount++
	}
	return vals
}

func (hm *HashMap[K, T]) Clear() {
	hm.m = make(map[K]T)
}
