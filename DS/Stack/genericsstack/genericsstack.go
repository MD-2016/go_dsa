package genericsstack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) Peek() (T, error) {
	var empty T
	if stack.IsEmpty() {
		return empty, errors.New("stack is empty. Cannot peek")
	}

	elm := stack.elements[len(stack.elements)-1]
	return elm, nil
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Pop() (T, error) {
	var empty T
	if stack.IsEmpty() {
		return empty, errors.New("Stack is empty. Cannot pop")
	}

	popped := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return popped, nil
}

func (stack *Stack[T]) ToString() string {
	return fmt.Sprintf("%v", stack.elements)
}

func (stack *Stack[T]) Reverse() []T {
	stackL := len(stack.elements)
	reverseStack := make([]T, stackL)
	for i, elm := range stack.elements {
		reverseStack[stackL-1-i] = elm
	}

	return reverseStack
}
