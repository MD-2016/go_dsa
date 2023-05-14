package bidmap

import (
	"fmt"
	"go_dsa/DS/hashmap"
)

type BidMap[K, T comparable] struct {
	toMap      hashmap.HashMap[K, T]
	inverseMap hashmap.HashMap[T, K]
}

func New[K, T comparable]() *BidMap[K, T] {
	return &BidMap[K, T]{*hashmap.New[K, T](), *hashmap.New[T, K]()}
}

func (bm *BidMap[K, T]) Put(key K, val T) {

	if vk, check := bm.toMap.Get(key); check {
		bm.inverseMap.Remove(vk)
	}
	if kv, check := bm.inverseMap.Get(val); check {
		bm.toMap.Remove(kv)
	}
	bm.toMap.Put(key, val)
	bm.inverseMap.Put(val, key)
}

func (bm *BidMap[K, T]) Get(key K) (T, bool) {
	var exists bool
	var val T
	val, exists = bm.toMap.Get(key)
	return val, exists
}

func (bm *BidMap[K, T]) GetKey(val T) (K, bool) {
	var exists bool
	var key K
	key, exists = bm.inverseMap.Get(val)
	return key, exists
}

func (bm *BidMap[K, T]) IsEmpty() bool {
	return bm.Size() == 0
}

func (bm *BidMap[K, T]) Size() int {
	return bm.toMap.Size()
}

func (bm *BidMap[K, T]) Keys() []K {
	return bm.toMap.Keys()
}

func (bm *BidMap[K, T]) Values() []T {
	return bm.inverseMap.Keys()
}

func (bm *BidMap[K, T]) Clear() {
	bm.toMap.Clear()
	bm.inverseMap.Clear()
}

func (bm *BidMap[K, T]) Remove(key K) (T, bool) {
	var val T
	var exists bool
	if val, exists := bm.toMap.Get(key); exists {
		bm.toMap.Remove(key)
		bm.inverseMap.Remove(val)
	}

	return val, exists
}

func (bm *BidMap[K, T]) ToString() string {
	mapStr := ""
	mapStr += fmt.Sprintf("%v", bm.toMap)
	return mapStr
}
