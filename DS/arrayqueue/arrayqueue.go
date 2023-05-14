package arrayqueue

import (
	"errors"
	"fmt"
	"go_dsa/DS/arraylist"
	"strings"
)

type Queue[T comparable] struct {
	list *arraylist.List[T]
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{list: arraylist.New[T]()}
}

func (queue *Queue[T]) Enqueue(val T) {
	queue.list.Add(val)
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	var val T
	var noError bool
	val, noError = queue.list.Get(0)
	if !noError {
		fmt.Sprintln(errors.New("unable to dequeue from queue"))
		return val, false
	} else {
		val, noError = queue.list.Remove(0)
		return val, noError
	}
}

func (queue *Queue[T]) Peek() (val T, errorFound bool) {
	return queue.list.Get(0)
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.IsEmpty()
}

func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

func (queue *Queue[T]) Clear() {
	queue.list.Clear()
}

func (queue *Queue[T]) Values() []T {
	return queue.list.Values()
}

func (queue *Queue[T]) ToString() string {
	str := ""
	vals := []string{}
	for _, val := range queue.list.Values() {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	str += strings.Join(vals, ", ")
	return str
}

func (queue *Queue[T]) InRange(i int) bool {
	return i >= 0 && i < queue.list.Size()
}
