package main

/*
Methods: functions attached to a type (usually a struct).
Value receiver  (r Type)  -> gets a COPY, can't mutate the original
Pointer receiver (r *Type) -> can mutate the original, avoids copying
*/

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Value receiver: read-only operation, no need to mutate
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver: mutates the original struct
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 4, Height: 5}
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())

	rect.Scale(2) // Go auto-takes the address: (&rect).Scale(2)
	fmt.Println("After Scale(2):", rect)

	c := Circle{Radius: 3}
	fmt.Println("Circle area:", c.Area())

	// Methods are just functions with a special first argument under the hood
	areaFunc := Rectangle.Area
	fmt.Println("Called via method value:", areaFunc(rect))
}
