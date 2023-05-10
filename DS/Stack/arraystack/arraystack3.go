package arraystack3

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func new[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) size() int {
	return len(stack.elements)
}

func (stack *Stack[T]) isEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) peek() (T, error) {
	var empty T
	if stack.isEmpty() {
		return empty, errors.New("Stack is empty. Cannot peek\n")
	}

	elm := stack.elements[len(stack.elements)-1]
	return elm, nil
}

func (stack *Stack[T]) push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) pop() (T, error) {
	var empty T
	if stack.isEmpty() {
		return empty, errors.New("Stack is empty. Cannot pop")
	}

	popped := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return popped, nil
}

func (stack *Stack[T]) toString() string {
	return fmt.Sprintf("%v", stack.elements)
}

func (stack *Stack[T]) reverse() []T {
	stackL := len(stack.elements)
	reverseStack := make([]T, stackL)
	for i, elm := range stack.elements {
		reverseStack[stackL-1-i] = elm
	}

	return reverseStack
}
