package linkedset

import (
	"go_dsa/DS/doublylinkedlist"
)

type LinkedSet[T comparable] struct {
	ordering *doublylinkedlist.DoublyLinkedList[T]
	set      map[T]struct{}
}

var elmExists = struct{}{}

func New[T comparable](vals ...T) *LinkedSet[T] {
	set := &LinkedSet[T]{
		set:      make(map[T]struct{}),
		ordering: doublylinkedlist.New[T](),
	}
	if len(vals) > 0 {
		set.Add(vals...)
	}
	return set
}

func (set *LinkedSet[T]) Add(elms ...T) {
	for _, elm := range elms {
		if _, contain := set.set[elm]; !contain {
			set.set[elm] = elmExists
			set.ordering.Append(elm)
		}
	}
}

func (set *LinkedSet[T]) Contains(elms ...T) bool {
	var found bool = false
	for _, elm := range elms {
		if _, contain := set.set[elm]; !contain {
			return found
		}
	}
	found = true
	return found
}

func (set *LinkedSet[T]) Remove(elms ...T) T {
	var removed T
	for _, elm := range elms {
		if _, contain := set.set[elm]; contain {
			delete(set.set, elm)
			i := set.ordering.IndexOf(elm)
			removed = set.ordering.Remove(i)
		}
	}

	return removed
}

func (set *LinkedSet[T]) IsEmpty() bool {
	return set.Size() == 0
}

func (set *LinkedSet[T]) Size() int {
	return set.ordering.Size()
}

func (set *LinkedSet[T]) Clear() {
	set.set = make(map[T]struct{})
	set.ordering.Clear()
}

func (set *LinkedSet[T]) Intersection(set2 *LinkedSet[T]) *LinkedSet[T] {
	resultingSet := New[T]()

	if set.Size() <= set2.Size() {
		for elm := range set.set {
			if _, contain := set2.set[elm]; contain {
				resultingSet.Add(elm)
			}
		}
	} else {
		for elm := range set.set {
			if _, contain := set.set[elm]; contain {
				resultingSet.Add(elm)
			}
		}
	}

	return resultingSet
}

func (set *LinkedSet[T]) Difference(set2 *LinkedSet[T]) *LinkedSet[T] {
	resultingSet := New[T]()

	for elm := range set.set {
		if _, contain := set2.set[elm]; !contain {
			resultingSet.Add(elm)
		}
	}

	return resultingSet
}

func (set *LinkedSet[T]) Union(set2 *LinkedSet[T]) *LinkedSet[T] {
	resultingSet := New[T]()

	for elm := range set.set {
		resultingSet.Add(elm)
	}
	for elm := range set2.set {
		resultingSet.Add(elm)
	}

	return resultingSet
}
