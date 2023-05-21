package binarysearchtree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	errDuplicate error = errors.New("bst: duplicate value found on tree")
	errNotFound  error = errors.New("bst: node not found")
)

/// Start of node
/// #####################################################

type Node[T constraints.Ordered] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func newNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (n *Node[T]) Value() T {
	return n.val
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func insertNode[T constraints.Ordered](n *Node[T], val T) (*Node[T], error) {
	if n == nil {
		return newNode(val), nil
	}

	if n.val == val {
		return nil, errDuplicate
	}

	if val > n.val {
		right, err := insertNode(n.right, val)

		if err != nil {
			return nil, err
		}

		n.right = right
	} else if val < n.val {
		left, err := insertNode(n.left, val)

		if err != nil {
			return nil, err
		}

		n.left = left
	}

	return n, nil
}

func removeNode[T constraints.Ordered](n *Node[T], val T) (*Node[T], error) {
	if n == nil {
		return nil, errNotFound
	}

	if val > n.val {
		right, err := removeNode(n.right, val)

		if err != nil {
			return nil, err
		}

		n.right = right
	} else if val < n.val {
		left, err := removeNode(n.left, val)
		if err != nil {
			return nil, err
		}

		n.left = left
	} else {
		if n.left != nil && n.right != nil {
			successor := leastNode(n.right)
			value := successor.val

			right, err := removeNode(n.right, value)
			if err != nil {
				return nil, err
			}

			n.right = right
			n.val = value
		} else if n.left != nil || n.right != nil {
			if n.left != nil {
				n = n.left
			} else {
				n = n.right
			}
		} else if n.left == nil && n.right == nil {
			n = nil
		}
	}

	return n, nil
}

func findNode[T constraints.Ordered](n *Node[T], val T) *Node[T] {
	if n == nil {
		return nil
	}

	if n.val == val {
		return n
	}

	if val > n.val {
		return findNode(n.right, val)
	}

	if val < n.val {
		return findNode(n.left, val)
	}

	return nil
}

func leastNode[T constraints.Ordered](n *Node[T]) *Node[T] {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n
	}

	return leastNode(n.left)
}

func walk[T constraints.Ordered](n *Node[T]) {
	if n == nil {
		return
	}

	walk(n.left)
	fmt.Println(n.val)
	walk(n.right)
}

/// end of node
/// #####################################################

type BST[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

func (bst *BST[T]) Insert(val T) error {
	root, err := insertNode(bst.root, val)

	if err != nil {
		return err
	}

	bst.root = root
	return nil
}

func (bst *BST[T]) Remove(val T) error {
	root, err := removeNode(bst.root, val)

	if err != nil {
		return err
	}

	bst.root = root
	return nil
}

func (bst *BST[T]) Find(val T) *Node[T] {
	return findNode(bst.root, val)
}

func (bst *BST[T]) Walk() {
	walk(bst.root)
}
