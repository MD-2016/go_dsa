package priorityqueue

import (
	"fmt"
	"go_dsa/DS/binaryheap"
	"go_dsa/DS/utils"
	"strings"
)

type PriorityQueue[T comparable] struct {
	bheap      *binaryheap.BinaryHeap[T]
	Comparator utils.Comparator[T]
}

func NewWithCompare[T comparable](compare utils.Comparator[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{bheap: binaryheap.New[T](compare), Comparator: compare}
}

func (pq *PriorityQueue[T]) Enqueue(val T) {
	pq.bheap.Push(val)
}

func (pq *PriorityQueue[T]) Dequeue() (val T, check bool) {
	return pq.bheap.Pop()
}

func (pq *PriorityQueue[T]) Peek() (val T, check bool) {
	return pq.bheap.Peek()
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.bheap.IsEmpty()
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.bheap.Size()
}

func (pq PriorityQueue[T]) Clear() {
	pq.bheap.Clear()
}

func (pq *PriorityQueue[T]) Values() []T {
	return pq.bheap.Values()
}

func (pq *PriorityQueue[T]) ToString() string {
	pqStr := ""
	vals := make([]string, pq.bheap.Size())
	for i, val := range pq.bheap.Values() {
		vals[i] = fmt.Sprintf("%v", val)
	}
	pqStr += strings.Join(vals, ", ")
	return pqStr
}
