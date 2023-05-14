package doublylinkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type DoublyLinkedList[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

type node[T comparable] struct {
	val  T
	prev *node[T]
	next *node[T]
}

func New[T comparable](vals ...T) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{}
	if len(vals) > 0 {
		list.Add(vals...)
	}
	return list
}

func (list *DoublyLinkedList[T]) Add(vals ...T) {
	for _, val := range vals {
		newNode := &node[T]{val: val, prev: list.tail}
		if list.length == 0 {
			list.head = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		list.length++
	}
}

func (list *DoublyLinkedList[T]) Append(vals ...T) {
	list.Add(vals...)
}

func (list *DoublyLinkedList[T]) Prepend(vals ...T) {
	for i := len(vals) - 1; i >= 0; i-- {
		newNode := &node[T]{val: vals[i], next: list.head}
		if list.length == 0 {
			list.head = newNode
			list.tail = newNode
		} else {
			list.head.prev = newNode
			list.head = newNode
		}
		list.length++
	}
}

func (list *DoublyLinkedList[T]) Get(i int) (T, bool) {
	if !list.InRange(i) {
		var empty T
		fmt.Sprintln(errors.New("out of range in doubly linked list"))
		return empty, false
	}

	if list.length-i < i {
		node := list.tail
		for n := list.length - 1; n != i; n, node = n-1, node.prev {
		}
		return node.val, true
	}
	node := list.head
	for n := 0; n != i; n, node = n+1, node.next {
	}
	return node.val, true
}

func (list *DoublyLinkedList[T]) Remove(i int) T {
	var empty T
	var removed T
	if !list.InRange(i) {
		fmt.Sprintln(errors.New("out of range in doubly linked list"))
		return empty
	}

	if list.length == 1 {
		list.Clear()
		return empty
	}

	var node *node[T]
	if list.length-i < i {
		node = list.tail
		for n := list.length - 1; n != i; n, node = n-1, node.prev {
		}
	} else {
		node = list.head
		for n := 0; n != i; n, node = n+1, node.next {
		}
	}

	if node == list.head {
		list.head = node.next
	}
	if node == list.tail {
		list.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	removed = node.val
	node = nil
	list.length--
	return removed
}

func (list *DoublyLinkedList[T]) Size() int {
	return list.length
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *DoublyLinkedList[T]) Contains(vals ...T) bool {
	if len(vals) == 0 {
		return true
	}
	if list.length == 0 {
		return false
	}
	for _, val := range vals {
		valFound := false
		for node := list.head; node != nil; node = node.next {
			if node.val == val {
				valFound = true
				break
			}
		}

		if !valFound {
			return false
		}
	}
	return true
}

func (list *DoublyLinkedList[T]) Values() []T {
	vals := make([]T, list.length)
	for i, node := 0, list.head; node != nil; i, node = i+1, node.next {
		vals[i] = node.val
	}
	return vals
}

func (list *DoublyLinkedList[T]) IndexOf(val T) int {
	if list.length == 0 {
		return -1
	}

	for i, node := range list.Values() {
		if node == val {
			return i
		}
	}
	return -1
}

func (list *DoublyLinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

func (list *DoublyLinkedList[T]) InRange(i int) bool {
	return i >= 0 && i < list.length
}

func (list *DoublyLinkedList[T]) Insert(i int, vals ...T) {

	if !list.InRange(i) {
		if i == list.length {
			list.Add(vals...)
		}
		return
	}

	list.length += len(vals)

	var previousNode *node[T]
	var nodeFound *node[T]

	if list.length-i < i {
		nodeFound = list.tail
		for n := list.length - 1; n != i; n, nodeFound = n-1, nodeFound.prev {
			previousNode = nodeFound.prev
		}
	} else {
		nodeFound = list.head
		for n := 0; n != i; n, nodeFound = n+1, nodeFound.next {
			previousNode = nodeFound
		}
	}

	if nodeFound == list.head {
		oldNode := list.head
		for i, val := range vals {
			newNode := &node[T]{val: val}
			if i == 0 {
				list.head = newNode
			} else {
				newNode.prev = previousNode
				previousNode.next = newNode
			}
			previousNode = newNode
		}
		oldNode.prev = previousNode
		previousNode.next = oldNode
	} else {
		oldNode := previousNode.next
		for _, val := range vals {
			newNode := &node[T]{val: val}
			newNode.prev = previousNode
			previousNode.next = newNode
			previousNode = newNode
		}
		oldNode.prev = previousNode
		previousNode.next = oldNode
	}
}

func (list *DoublyLinkedList[T]) ToString() string {
	str := ""
	vals := []string{}
	for i := list.head; i != nil; i = i.next {
		vals = append(vals, fmt.Sprintf("%v", i.val))
	}
	str += strings.Join(vals, ", ")
	return str
}
