package utils

import "time"

type CompareNum interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Comparator[T comparable] func(a, b T) int

func StringCompare(a, b string) int {
	s1 := a
	s2 := b
	min := len(s2)

	if len(s1) < len(s2) {
		min = len(s1)
	}

	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

func NumberCompare[T CompareNum](a, b T) int {
	n1 := a
	n2 := b
	switch {
	case n1 > n2:
		return 1
	case n1 < n2:
		return -1
	default:
		return 0
	}
}

func ByteCompare(a, b byte) int {
	b1 := a
	b2 := b
	switch {
	case b1 > b2:
		return 1
	case b1 < b2:
		return -1
	default:
		return 0
	}
}

func RuneCompare(a, b rune) int {
	r1 := a
	r2 := b
	switch {
	case r1 > r2:
		return 1
	case r1 < r2:
		return -1
	default:
		return 0
	}
}

func TimeCompare(a, b time.Time) int {
	t1 := a
	t2 := b
	switch {
	case t1.After(t2):
		return 1
	case t1.Before(t2):
		return -1
	default:
		return 0
	}
}
