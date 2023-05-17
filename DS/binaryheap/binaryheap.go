package binaryheap

import (
	"fmt"
	"go_dsa/DS/arraylist"
	"go_dsa/DS/utils"
	"strings"
)

type BinaryHeap[T comparable] struct {
	list    *arraylist.List[T]
	compare utils.Comparator[T]
}

func New[T comparable](compare utils.Comparator[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{list: arraylist.New[T](), compare: compare}
}

func NewNumCom[T utils.CompareNum]() *BinaryHeap[T] {
	return &BinaryHeap[T]{list: arraylist.New[T](), compare: utils.NumberCompare[T]}
}

func NewStrCom() *BinaryHeap[string] {
	return &BinaryHeap[string]{list: arraylist.New[string](), compare: utils.StringCompare}
}

func (bh *BinaryHeap[T]) Push(vals ...T) {
	if len(vals) == 1 {
		bh.list.Add(vals[0])
		bh.heapUp()
	} else {
		for _, val := range vals {
			bh.list.Add(val)
		}
		length := bh.list.Size()/2 + 1
		for i := length; i >= 0; i-- {
			bh.heapdown()
		}
	}
}

func (bh *BinaryHeap[T]) heapUp() {
	i := bh.list.Size() - 1
	for parentIn := (i - 1) >> 1; i > 0; parentIn = (i - 1) >> 1 {
		inVal, _ := bh.list.Get(i)
		parentVal, _ := bh.list.Get(parentIn)
		if bh.compare(parentVal, inVal) <= 0 {
			break
		}
		bh.list.Swap(i, parentIn)
		i = parentIn
	}
}

func (bh *BinaryHeap[T]) heapdownindex(i int) {
	length := bh.list.Size()
	for leftIn := i<<1 + 1; leftIn < length; leftIn = i<<1 + 1 {
		rightIn := i<<1 + 2
		smallIn := leftIn
		leftVal, _ := bh.list.Get(leftIn)
		rightVal, _ := bh.list.Get(rightIn)
		if rightIn < length && bh.compare(leftVal, rightVal) > 0 {
			smallIn = rightIn
		}
		inVal, _ := bh.list.Get(i)
		smallVal, _ := bh.list.Get(smallIn)
		if bh.compare(inVal, smallVal) > 0 {
			bh.list.Swap(i, smallIn)
		} else {
			break
		}
		i = smallIn
	}
}

func (bh *BinaryHeap[T]) heapdown() {
	bh.heapdownindex(0)
}

func (bh *BinaryHeap[T]) Pop() (val T, check bool) {
	val, check = bh.list.Get(0)
	if !check {
		return
	}
	lastIn := bh.list.Size() - 1
	bh.list.Swap(0, lastIn)
	bh.list.Remove(lastIn)
	bh.heapdown()
	return
}

func (bh *BinaryHeap[T]) Peek() (val T, check bool) {
	return bh.list.Get(0)
}

func (bh *BinaryHeap[T]) IsEmpty() bool {
	return bh.list.IsEmpty()
}

func (bh *BinaryHeap[T]) Size() int {
	return bh.list.Size()
}

func (bh *BinaryHeap[T]) Clear() {
	bh.list.Clear()
}

func (bh *BinaryHeap[T]) Values() []T {
	vals := make([]T, bh.list.Size())
	for itr := bh.Iterator(); itr.Next(); {
		vals[itr.Index()] = itr.Value()
	}
	return vals
}

func (bh *BinaryHeap[T]) ToString() string {
	bhStr := ""
	vals := []string{}
	for itr := bh.Iterator(); itr.Next(); {
		vals = append(vals, fmt.Sprintf("%v", itr.Value()))
	}
	bhStr += strings.Join(vals, ", ")
	return bhStr
}

func (bh *BinaryHeap[T]) InRange(index int) bool {
	return index >= 0 && index < bh.list.Size()
}
