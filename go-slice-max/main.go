package main

import (
	"cmp"
	"fmt"
)

type Stack[T cmp.Ordered] struct {
	items []T
	max   []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)

	size := len(s.max)
	if size == 0 {
		s.max = append(s.max, value)
		return
	}

	m := s.max[size-1]
	if m <= value {
		s.max = append(s.max, value)
		return
	}

	s.max = append(s.max, m)
}

func (s *Stack[T]) Pop() *T {
	size := len(s.items)
	if size == 0 {
		var empty T
		return &empty
	}

	item := s.items[size-1]

	s.items = s.items[:size-1]
	s.max = s.max[:size-1]

	return &item
}

func (s *Stack[T]) Max() *T {
	size := len(s.max)
	if size == 0 {
		var empty T
		return &empty
	}

	return &s.max[size-1]
}

func main() {
	var stack Stack[int]

	stack.Push(10)
	stack.Push(2)
	stack.Push(5)
	stack.Push(-1)
	stack.Push(7)

	fmt.Println(*stack.Max())
	fmt.Println(*stack.Pop())
	fmt.Println()

	fmt.Println(*stack.Max())
	fmt.Println(*stack.Pop())
	fmt.Println()

	fmt.Println(*stack.Max())
	fmt.Println(*stack.Pop())
	fmt.Println()

	fmt.Printf("%+v", stack)
}
