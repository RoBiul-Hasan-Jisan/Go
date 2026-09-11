package main

/*
Generics (Go 1.18+): write functions/types that work with any type
satisfying a constraint, without losing type safety or writing interface{}
boilerplate.
*/

import "fmt"

// Generic function: T can be any type that supports ordering (via `cmp.Ordered`-like constraint)
type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic type: a stack that works for any element type
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// Generic Map/Filter helpers - staples of functional-style Go
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	fmt.Println("Sum ints:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("Sum floats:", Sum([]float64{1.1, 2.2, 3.3}))
	fmt.Println("Max:", Max(10, 20))

	var s Stack[string]
	s.Push("a")
	s.Push("b")
	s.Push("c")
	top, _ := s.Pop()
	fmt.Println("Popped:", top)

	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	strs := Map(nums, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Println("As strings:", strs)
}
