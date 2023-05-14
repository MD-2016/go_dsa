package hashset

import (
	"fmt"
	"strings"
)

type HashSet[T comparable] struct {
	items map[T]struct{}
}

var itemDuplicate = struct{}{}

func New[T comparable](vals ...T) *HashSet[T] {
	hashset := &HashSet[T]{items: make(map[T]struct{})}
	if len(vals) > 0 {
		hashset.Add(vals...)
	}
	return hashset
}

func (set *HashSet[T]) Add(vals ...T) {
	for _, val := range vals {
		set.items[val] = itemDuplicate
	}
}

func (set *HashSet[T]) Remove(vals ...T) {
	for _, val := range vals {
		delete(set.items, val)
	}
}

func (set *HashSet[T]) Contains(vals ...T) bool {
	for _, val := range vals {
		if _, contain := set.items[val]; !contain {
			return false
		}
	}
	return true
}

func (set *HashSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

func (set *HashSet[T]) Size() int {
	return len(set.items)
}

func (set *HashSet[T]) Clear() {
	set.items = make(map[T]struct{})
}

func (set *HashSet[T]) ToString() string {
	str := ""
	elms := []string{}
	for i := range set.items {
		elms = append(elms, fmt.Sprintf("%v", i))
	}
	str += strings.Join(elms, ", ")
	return str
}

func (set *HashSet[T]) Intersection(set2 *HashSet[T]) *HashSet[T] {
	resultingSet := New[T]()

	if set.Size() <= set2.Size() {
		for elm := range set.items {
			if _, contain := set2.items[elm]; contain {
				resultingSet.Add(elm)
			}
		}
	} else {
		for elm := range set2.items {
			if _, contain := set.items[elm]; contain {
				resultingSet.Add(elm)
			}
		}
	}
	return resultingSet
}

func (set *HashSet[T]) Union(set2 *HashSet[T]) *HashSet[T] {
	resultingSet := New[T]()

	for elm := range set.items {
		resultingSet.Add(elm)
	}
	for elm := range set.items {
		resultingSet.Add(elm)
	}

	return resultingSet
}

func (set *HashSet[T]) Difference(set2 *HashSet[T]) *HashSet[T] {
	resultingSet := New[T]()

	for elm := range set.items {
		if _, contain := set2.items[elm]; !contain {
			resultingSet.Add(elm)
		}
	}

	return resultingSet
}
