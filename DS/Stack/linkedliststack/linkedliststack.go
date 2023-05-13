package linkedliststack

import "go_dsa/DS/singlylinkedlist"

type Stack[T comparable] struct {
	list *singlylinkedlist.List[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: &singlylinkedlist}
}
