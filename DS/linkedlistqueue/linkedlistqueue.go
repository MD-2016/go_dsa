package linkedlistqueue

import (
	"errors"
	"fmt"
	"go_dsa/DS/singlylinkedlist"
	"strings"
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

func (queue *Queue[T]) Dequeue() (T, error) {
	var val T
	var errorFound error
	val, errorFound = queue.list.Get(0)
	if errorFound != nil {
		val, errorFound = queue.list.Remove(0)
		return val, nil
	} else {
		return val, errors.New("queue is empty")
	}

}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

func (queue *Queue[T]) Clear() {
	queue.list.Clear()
}

func (queue *Queue[T]) ToString() string {
	str := "QueueVals\n"
	vals := []string{}
	for _, val := range queue.list.Values() {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	str += strings.Join(vals, ", ")
	return str
}
