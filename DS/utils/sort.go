package utils

import "sort"

func Sort[T comparable](vals []T, compare Comparator[T]) {
	sort.Sort(sortable[T]{vals, compare})
}

type sortable[T comparable] struct {
	vals    []T
	compare Comparator[T]
}

func (s sortable[T]) Len() int {
	return len(s.vals)
}

func (s sortable[T]) Swap(i, j int) {
	s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
}

func (s sortable[T]) Less(i, j int) bool {
	return s.compare(s.vals[i], s.vals[j]) < 0
}
