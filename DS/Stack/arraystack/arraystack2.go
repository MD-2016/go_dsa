package arraystack2

import (
	"errors"
)

type Stack struct {
	elements []interface{}
}

func New(elm ...interface{}) *Stack {
	return &Stack{
		elements: elm,
	}
}

func (stack *Stack) size() int {
	return len(stack.elements)
}

func (stack *Stack) isEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack) push(elm ...interface{}) {
	stack.elements = append(stack.elements, elm)
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("Stack is Empty")
	}

	popped := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return popped, nil
}
