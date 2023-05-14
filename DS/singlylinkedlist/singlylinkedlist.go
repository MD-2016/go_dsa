package singlylinkedlist

import (
	"errors"
	"fmt"
)

type List[T comparable] struct {
	first  *node[T]
	last   *node[T]
	length int
}

type node[T comparable] struct {
	next  *node[T]
	value T
}

func New[T comparable](vals ...T) *List[T] {
	newList := &List[T]{}
	if len(vals) > 0 {
		newList.Add(vals...)
	}
	return newList
}

func (list *List[T]) Add(vals ...T) {
	for _, val := range vals {
		newNode := &node[T]{value: val}
		if list.length == 0 {
			list.first = newNode
			list.last = newNode
		} else {
			list.last.next = newNode
			list.last = newNode
		}
		list.length++
	}
}

func (list *List[T]) Append(vals ...T) {
	list.Add(vals...)
}

func (list *List[T]) Prepend(vals ...T) {
	for val := len(vals) - 1; val >= 0; val-- {
		newNode := &node[T]{value: vals[val], next: list.first}
		list.first = newNode
		if list.length == 0 {
			list.last = newNode
		}
		list.length++
	}
}

func (list *List[T]) inRange(index int) bool {
	return index >= 0 && index < list.length
}

func (list *List[T]) Get(index int) (T, error) {
	if !list.inRange(index) {
		var empty T
		return empty, errors.New("List index is out of range")
	}

	currElm := list.first
	for i := 0; i != index; i, currElm = i+1, currElm.next {

	}

	return currElm.value, nil
}

func (list *List[T]) Remove(index int) (T, error) {
	var empty T
	if !list.inRange(index) {
		return empty, errors.New("index is not in range of list")
	}

	if list.length == 1 {
		list.Clear()
		return empty, nil
	}

	var previousNode *node[T]
	removedNode := list.first
	for i := 0; i != index; i, removedNode = i+1, removedNode.next {
		previousNode = removedNode
	}

	if removedNode == list.first {
		list.first = removedNode.next
	}

	if removedNode == list.last {
		list.last = previousNode
	}

	if previousNode != nil {
		previousNode.next = removedNode.next
	}

	empty = removedNode.value
	removedNode = nil
	list.length--
	return empty, nil
}

func (list *List[T]) Contains(vals ...T) bool {

	if len(vals) == 0 {
		return true
	}

	if list.length == 0 {
		return false
	}

	for _, val := range vals {
		isFound := false
		for foundNode := list.first; foundNode != nil; foundNode = foundNode.next {
			if foundNode.value == val {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}

func (list *List[T]) Size() int {
	return list.length
}

func (list *List[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *List[T]) Clear() {
	list.length = 0
	list.first = nil
	list.last = nil
}

func (list *List[T]) IndexOf(val T) int {
	if list.length == 0 {
		return -1
	}

	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		if node.value == val {
			return i
		}
	}
	return -1
}

func (list *List[T]) Insert(index int, vals ...T) {
	if !list.inRange(index) {
		if index == list.length {
			list.Add(vals...)
		}
		return
	}

	list.length += len(vals)

	var previousNode *node[T]
	currNode := list.first
	for i := 0; i != index; i, currNode = i+1, currNode.next {
		previousNode = currNode
	}

	if currNode == list.first {
		oldNext := list.first
		for i, val := range vals {
			newNode := &node[T]{value: val}
			if i == 0 {
				list.first = newNode
			} else {
				previousNode.next = newNode
			}
		}
		previousNode.next = oldNext
	} else {
		oldNext := previousNode.next
		for _, val := range vals {
			newNode := &node[T]{value: val}
			previousNode.next = newNode
			previousNode = newNode
		}
		previousNode.next = oldNext
	}
}

func (list *List[T]) ToString() []string {
	vals := []string{}
	for node := list.first; node != nil; node = node.next {
		vals = append(vals, fmt.Sprintf("%v", node.value))
	}
	return vals
}

func (list *List[T]) Values() []T {
	vals := make([]T, list.length, list.length)
	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		vals[i] = node.value
	}
	return vals
}
