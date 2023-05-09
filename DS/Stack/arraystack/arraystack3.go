package arraystack3

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) size() int {
	return len(stack.elements)
}
