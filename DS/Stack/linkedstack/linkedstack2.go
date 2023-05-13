package linkedstack2

type node[T any] struct {
	item T
	next *node[T]
}

type Stack[T any] struct {
	head   *node[T]
	length int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) size() int {
}
