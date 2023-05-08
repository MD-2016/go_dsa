package arraystack

import (
	"errors"
)

type Stack []interface{}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) size() int {
	return len(*stack)
}

func (stack *Stack) peek() (element interface{}) {
	element = (*stack)[len(*stack)-1]
	return element
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("stack is empty")
	}

	rest := len(*stack) - 1
	element := (*stack)[rest]
	(*stack) = (*stack)[:rest]
	return element, nil
}

func (stack *Stack) push(element interface{}) {
	(*stack) = append((*stack), element)
}
