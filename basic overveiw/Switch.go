package main

/*
Switch: cleaner alternative to long if-else chains.
No "fallthrough" by default (unlike C/Java) - each case breaks automatically.
*/

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4, 5: // multiple values in one case
		fmt.Println("Thursday or Friday")
	default:
		fmt.Println("Weekend")
	}

	// Switch without a condition == cleaner if-else chain
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 75:
		fmt.Println("Grade: B")
	case score >= 60:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// fallthrough: explicitly continue to next case
	num := 1
	switch num {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two (reached via fallthrough)")
	case 3:
		fmt.Println("three")
	}

	// Type switch (also shown in Interfaces.go)
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}
