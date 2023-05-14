package linkedlistqueue

import (
	"go_dsa/DS/singlylinkedlist"
)

type Queue[T comparable] struct {
	list *singlylinkedlist.List[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{list: &singlylinkedlist.List[T]{}}
}

func (queue *Queue[T]) Enqueue(val T) {
	queue.list.Add(val)
}

func (queue *Queue[T]) Peek() (T, error) {
	return queue.list.Get(0)
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func
