package arraystack2

import (
	"errors"
	"fmt"
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

func (stack *Stack) peek() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("Stack is empty no way to peek")
	}
	elm := stack.elements[len(stack.elements)-1]
	return elm, nil
}

// assumes the stack isn't empty
func (stack *Stack) reverse() []interface{} {
	stackLen := len(stack.elements)
	reversedStack := make([]interface{}, stackLen)
	for i, elm := range stack.elements {
		reversedStack[stackLen-1-i] = elm
	}
	return reversedStack
}

func (stack *Stack) toString() string {
	return fmt.Sprintf("%v", stack.elements...)
}
