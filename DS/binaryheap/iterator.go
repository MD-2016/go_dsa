package binaryheap

type Iterator[T comparable] struct {
	bh    *BinaryHeap[T]
	index int
}

func (bh *BinaryHeap[T]) Iterator() Iterator[T] {
	return Iterator[T]{bh: bh, index: -1}
}

func (iter *Iterator[T]) Next() bool {
	if iter.index < iter.bh.Size() {
		iter.index++
	}
	return iter.bh.InRange(iter.index)
}

func (iter *Iterator[T]) Value() T {
	begin, end := bitsInRange(iter.index)
	if end > iter.bh.Size() {
		end = iter.bh.Size()
	}
	tmp := New(iter.bh.compare)
	for i := begin; i < end; i++ {
		val, _ := iter.bh.list.Get(i)
		tmp.Push(val)
	}
	for i := 0; i < iter.index-begin; i++ {
		tmp.Pop()
	}
	val, _ := tmp.Pop()
	return val
}

func (iter *Iterator[T]) Index() int {
	return iter.index
}

func (iter *Iterator[T]) Begin() {
	iter.index = -1
}

func (iter *Iterator[T]) End() {
	iter.index = iter.bh.Size()
}

func bitsNum(b int) uint {
	var count uint
	for b != 0 {
		count++
		b >>= 1
	}
	return count
}

func bitsInRange(i int) (int, int) {
	bts := bitsNum(i+1) - 1
	begin := 1<<bts - 1
	end := begin + 1<<bts
	return begin, end
}
