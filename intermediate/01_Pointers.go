package main

/*
Pointers hold the MEMORY ADDRESS of a value.
&x  -> "address of x"
*p  -> "value pointed to by p" (dereference)
Go has pointers but NO pointer arithmetic (unlike C) - much safer.
*/

import "fmt"

func increment(n *int) {
	*n = *n + 1 // dereference and mutate the original int
}

func newCounter() *int {
	x := 0
	return &x // safe in Go: escape analysis moves x to the heap automatically
}

type Point struct{ X, Y int }

func movePoint(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	x := 10
	p := &x // p is a *int pointing at x

	fmt.Println("Value of x:", x)
	fmt.Println("Address stored in p:", p)
	fmt.Println("Value pointed to by p:", *p)

	*p = 20 // changes x through the pointer
	fmt.Println("x after *p = 20:", x)

	increment(&x)
	fmt.Println("x after increment():", x)

	counter := newCounter()
	*counter++
	fmt.Println("Counter via returned pointer:", *counter)

	pt := Point{X: 1, Y: 1}
	movePoint(&pt, 5, 5)
	fmt.Println("Point after move:", pt)

	// nil pointers - always check before dereferencing
	var np *int
	fmt.Println("Nil pointer:", np, "is nil?", np == nil)

	// new() allocates and returns a pointer to a zero value
	num := new(int)
	*num = 42
	fmt.Println("new(int) ->", *num)
}
