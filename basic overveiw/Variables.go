package main

/*
Go Variables Cheat-Sheet
------------------------
var name type = value      // explicit type
var name = value           // type inferred
name := value              // short declaration (inside functions only)
const name = value         // compile-time constant
Multiple vars: var a, b, c int  OR  a, b := 1, "two"
Zero values: int=0, float64=0, string="", bool=false, pointer=nil
*/

import "fmt"

func main() {
	// 1. Explicit type
	var age int = 25

	// 2. Type inferred from value
	var name = "Gopher"

	// 3. Short declaration (most common inside functions)
	city := "Dhaka"

	// 4. Multiple variables at once
	var x, y, z = 1, 2.5, "three"

	// 5. Zero values (declared but not assigned)
	var count int
	var price float64
	var isActive bool
	var label string

	// 6. Constants
	const Pi = 3.14159
	const AppName = "GoLearner"

	fmt.Println("Name:", name, "| Age:", age, "| City:", city)
	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Println("Zero values -> count:", count, "price:", price, "isActive:", isActive, "label:", label)
	fmt.Println("Constants -> Pi:", Pi, "AppName:", AppName)

	// 7. Type conversion (Go never converts types implicitly)
	var whole int = 10
	var decimal float64 = float64(whole) / 3
	fmt.Println("Converted division:", decimal)
}
