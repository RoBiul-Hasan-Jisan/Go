package main

/*
Closures: a function that "closes over" variables from its surrounding scope.
The returned function keeps a live reference to those variables, not a copy.
*/

import "fmt"

// counter() returns a closure that remembers `count` between calls
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeMultiplier demonstrates closures capturing a parameter
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	next := counter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	// A second closure has its OWN independent state
	next2 := counter()
	fmt.Println("next2:", next2()) // 1 (independent from `next`)
	fmt.Println("next:", next())   // 4

	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("double(5):", double(5))
	fmt.Println("triple(5):", triple(5))

	// Classic Go gotcha: capturing a loop variable (fixed since Go 1.22,
	// each iteration gets its own `i`, but it's good to know the pattern)
	funcs := make([]func(), 0)
	for i := 0; i < 3; i++ {
		i := i // pre-1.22 idiom: shadow to capture a fresh copy per iteration
		funcs = append(funcs, func() { fmt.Println("captured i:", i) })
	}
	for _, f := range funcs {
		f()
	}
}
