package arraylist

import (
	"errors"
	"fmt"
	"strings"
)

type List[T comparable] struct {
	items  []T
	length int
}

const (
	growFac   = int(2)
	shrinkFac = float32(0.25)
)

func New[T comparable](vals ...T) *List[T] {
	newList := &List[T]{}
	if len(vals) > 0 {
		newList.Add(vals...)
	}
	return newList
}

func (list *List[T]) Add(vals ...T) {
	list.Grow(len(vals))
	for _, val := range vals {
		list.items[list.length] = val
		list.length++
	}
}

func (list *List[T]) Get(index int) (T, bool) {
	if !list.InRange(index) {
		var empty T
		fmt.Sprintln(errors.New("not in range of list"))
		return empty, false
	}

	return list.items[index], true
}

func (list *List[T]) Contains(vals ...T) bool {
	for _, val := range vals {
		isFound := false
		for i := 0; i < list.length; i++ {
			if list.items[i] == val {
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

func (list *List[T]) Values() []T {
	newItems := make([]T, list.length)
	copy(newItems, list.items[:list.length])
	return newItems
}

func (list *List[T]) IndexOf(val T) int {
	if list.length == 0 {
		return -1
	}

	for i, elm := range list.items {
		if elm == val {
			return i
		}
	}
	return -1
}

func (list *List[T]) IsEmpty() bool {
	return list.length == 0
}

func (list *List[T]) Size() int {
	return list.length
}

func (list *List[T]) Clear() {
	list.length = 0
	list.items = []T{}
}

func (list *List[T]) Swap(index1, index2 int) {
	if list.InRange(index1) && list.InRange(index2) {
		list.items[index1], list.items[index2] = list.items[index2], list.items[index1]
	}
}

func (list *List[T]) Insert(i int, vals ...T) {
	if !list.InRange(i) {
		if i == list.length {
			list.Add(vals...)
		}
		return
	}

	newSize := len(vals)
	list.Grow(newSize)
	list.length += 1
	copy(list.items[i+1:], list.items[i:list.length-newSize])
	copy(list.items[i:], vals)
}

func (list *List[T]) Remove(index int) (T, bool) {
	if !list.InRange(index) {
		var empty T
		fmt.Sprintln(errors.New("not in range of list"))
		return empty, false
	}

	var element T
	list.items[index] = element
	copy(list.items[index:], list.items[index+1:list.length])
	list.length--

	list.Shrink()
	return element, true
}

func (list *List[T]) Grow(growthVal int) {
	currCap := cap(list.items)
	if list.length+growthVal >= currCap {
		newCap := int(float32(growFac) * float32(currCap+growthVal))
		list.Resize(newCap)
	}
}

func (list *List[T]) Shrink() {
	if shrinkFac == 0.0 {
		return
	}

	currCap := cap(list.items)
	if list.length <= int(float32(currCap)*shrinkFac) {
		list.Resize(list.length)
	}
}

func (list *List[T]) Resize(newCap int) {
	newElm := make([]T, newCap)
	copy(newElm, list.items)
	list.items = newElm
}

func (list *List[T]) InRange(index int) bool {
	return index >= 0 && index < list.length
}

func (list *List[T]) ToString() string {
	str := ""
	vals := []string{}
	for _, val := range list.items[:list.length] {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	str += strings.Join(vals, ", ")
	return str
}
