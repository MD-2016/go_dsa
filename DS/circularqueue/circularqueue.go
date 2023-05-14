package circularqueue

import (
	"fmt"
	"strings"
)

type CircularQueue[T comparable] struct {
	vals      []T
	begin     int
	end       int
	isFull    bool
	maxLength int
	length    int
}

func New[T comparable](maxLength int) *CircularQueue[T] {
	if maxLength < 1 {
		panic("must have 1 or more size")
	}
	queue := &CircularQueue[T]{maxLength: maxLength}
	queue.Clear()
	return queue
}

func (queue *CircularQueue[T]) Clear() {
	queue.vals = make([]T, queue.maxLength)
	queue.begin = 0
	queue.end = 0
	queue.isFull = false
	queue.length = 0
}

func (queue *CircularQueue[T]) Enqueue(val T) {
	if queue.Full() {
		queue.Dequeue()
	}
	queue.vals[queue.end] = val
	queue.end = queue.end + 1
	if queue.end >= queue.maxLength {
		queue.end = 0
	}
	if queue.end == queue.begin {
		queue.isFull = true
	}

	queue.length = queue.CheckSize()
}

func (queue *CircularQueue[T]) Dequeue() (val T, errorFound bool) {
	var empty T

	if queue.IsEmpty() {
		errorFound = false
		return empty, errorFound
	}

	val, errorFound = queue.vals[queue.begin], true

	if val != empty {
		queue.vals[queue.begin] = empty
		queue.begin = queue.begin + 1
		if queue.begin >= queue.maxLength {
			queue.begin = 0
		}
		queue.isFull = false
	}

	queue.length = queue.length - 1

	return val, errorFound
}

func (queue *CircularQueue[T]) Peek() (val T, errorFound bool) {
	if queue.IsEmpty() {
		var empty T
		errorFound = false
		return empty, errorFound
	}

	errorFound = true
	return queue.vals[queue.begin], errorFound
}

func (queue *CircularQueue[T]) CheckSize() int {
	if queue.end < queue.begin {
		return queue.maxLength - queue.begin + queue.end
	} else if queue.end == queue.begin {
		if queue.isFull {
			return queue.maxLength
		}
		return 0
	}
	return queue.end - queue.begin
}

func (queue *CircularQueue[T]) IsEmpty() bool {
	return queue.length == 0
}

func (queue *CircularQueue[T]) Full() bool {
	return queue.Size() == queue.maxLength
}

func (queue *CircularQueue[T]) Size() int {
	return queue.length
}

func (queue *CircularQueue[T]) Values() []T {
	vals := make([]T, queue.Size())
	for i := 0; i < queue.Size(); i++ {
		vals[i] = queue.vals[(queue.begin+i)%queue.maxLength]
	}
	return vals
}

func (queue *CircularQueue[T]) InRange(i int) bool {
	return i >= 0 && i < queue.length
}

func (queue *CircularQueue[T]) ToString() string {
	str := ""
	vals := []string{}
	for _, val := range queue.Values() {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	str += strings.Join(vals, ", ")
	return str
}
