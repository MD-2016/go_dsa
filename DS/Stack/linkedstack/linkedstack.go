package linkedstack

import "errors"

type node struct {
	items interface{}
	next  *node
}

type Stack struct {
	head   *node
	length int
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) size() int {
	return stack.length
}

func (stack *Stack) isEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) peek() (interface{}, error) {
	if stack.length == 0 {
		return nil, errors.New("Stack is empty. Can't peek!")
	}

	return stack.head.items, nil
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.length == 0 {
		return nil, errors.New("Stack is empty. Can't pop!")
	}
	item := stack.head.items
	nextItem := stack.head.next
	stack.head.next = nil
	stack.head = nextItem
	stack.length--
	return item, nil
}

func (stack *Stack) push(item interface{}) {
	node := &node{
		items: item,
	}

	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
	stack.length++
}
