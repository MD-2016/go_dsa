package deque

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type Deque[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T comparable]() *Deque[T] {
	return &Deque[T]{}
}

func (dq *Deque[T]) AddFirst(val T) {
	newNode := new(Node[T])
	newNode.value = val
	if dq.length == 0 {
		dq.head = newNode
		dq.tail = newNode
	} else {
		newNode.next = dq.head
		dq.head.prev = newNode
		dq.head = newNode
	}
	dq.length++
}

func (dq *Deque[T]) RemoveFirst() (T, error) {
	if dq.length == 0 {
		return *new(T), errors.New("deque is empty. Can't remove from front")
	}
	val := dq.head.value
	if dq.length == 1 {
		dq.head = nil
		dq.tail = nil
	} else {
		dq.head = dq.head.next
		dq.head.prev = nil
	}
	dq.length--
	return val, nil
}

func (dq *Deque[T]) GetFirst() (T, error) {
	if dq.length == 0 {
		return *new(T), errors.New("deque is empty can't get first")
	}
	return dq.head.value, nil
}

func (dq *Deque[T]) AddLast(val T) {
	newNode := new(Node[T])
	newNode.value = val
	if dq.length == 0 {
		dq.head = newNode
		dq.tail = newNode
	} else {
		newNode.prev = dq.tail
		dq.tail.next = newNode
		dq.tail = newNode
	}
	dq.length++
}

func (dq *Deque[T]) RemoveLast() (T, error) {
	if dq.length == 0 {
		return *new(T), errors.New("deque is empty. Cannot remove from end")
	}
	val := dq.tail.value
	if dq.length == 1 {
		dq.head = nil
		dq.tail = nil
	} else {
		dq.tail = dq.tail.prev
		dq.tail.next = nil
	}
	dq.length--
	return val, nil
}

func (dq *Deque[T]) GetLast() (T, error) {
	if dq.length == 0 {
		return *new(T), errors.New("deque is empty. Cannot get last")
	}
	return dq.tail.value, nil
}

func (dq *Deque[T]) Size() int {
	return dq.length
}

func (dq *Deque[T]) IsEmpty() bool {
	return dq.length == 0
}

func (dq *Deque[T]) Values() []T {
	vals := make([]T, dq.Size())
	for i, node := 0, dq.head; node != nil; i, node = i+1, node.next {
		vals[i] = node.value
	}
	return vals
}

func (dq *Deque[T]) ToString() string {
	dqStr := ""
	vals := make([]string, dq.Size())
	for i, val := range dq.Values() {
		vals[i] = fmt.Sprintf("%v", val)
	}
	dqStr += strings.Join(vals, ", ")
	return dqStr
}
