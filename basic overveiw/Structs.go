package main

/*
Structs: custom composite types that group related fields together.
*/

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address // nested struct (composition)
}

func main() {
	// Create with field names
	p1 := Person{
		Name: "Rafi",
		Age:  22,
		Address: Address{
			City:    "Chattogram",
			Country: "Bangladesh",
		},
	}
	fmt.Println("p1:", p1)
	fmt.Println("p1 city:", p1.Address.City)

	// Create positionally (order matters, less readable)
	p2 := Person{"Mina", 28, Address{"Dhaka", "Bangladesh"}}
	fmt.Println("p2:", p2)

	// Zero-value struct
	var p3 Person
	fmt.Println("Zero-value struct:", p3)

	// Modify fields
	p1.Age = 23
	fmt.Println("Updated age:", p1.Age)

	// Pointer to struct (common for mutation in functions)
	updateAge(&p1, 30)
	fmt.Println("After updateAge:", p1.Age)

	// Anonymous struct (quick, one-off use)
	point := struct{ X, Y int }{X: 5, Y: 10}
	fmt.Println("Anonymous struct:", point)

	// Struct comparison (works if all fields are comparable)
	a := Address{"Dhaka", "Bangladesh"}
	b := Address{"Dhaka", "Bangladesh"}
	fmt.Println("a == b?", a == b)
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge // no need to dereference explicitly, Go does it for you
}
