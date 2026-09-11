package main

/*
Interfaces: define a set of method signatures.
Any type that implements ALL those methods automatically satisfies the interface
(no "implements" keyword needed - it's implicit/structural typing).
*/

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return 3.14159 * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * 3.14159 * c.Radius }

// Function that accepts ANY type satisfying the Shape interface
func describe(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// The empty interface `any` (alias for interface{}) can hold a value of ANY type
func printAnything(v any) {
	fmt.Println("Value:", v)
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
	}

	for _, s := range shapes {
		describe(s)
	}

	// Type assertion: check the concrete type behind an interface value
	var s Shape = Circle{Radius: 2}
	if c, ok := s.(Circle); ok {
		fmt.Println("It's a circle with radius:", c.Radius)
	}

	// Type switch: branch on the concrete type
	for _, s := range shapes {
		switch v := s.(type) {
		case Rectangle:
			fmt.Println("Rectangle with width", v.Width)
		case Circle:
			fmt.Println("Circle with radius", v.Radius)
		default:
			fmt.Println("Unknown shape")
		}
	}

	printAnything(42)
	printAnything("hello")
	printAnything(true)
}
