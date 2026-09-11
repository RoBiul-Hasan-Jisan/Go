package main

/*
Functions in Go:
- Can return multiple values
- Support named return values
- Support variadic parameters (...)
- Are first-class values (can be passed around, assigned, returned)
*/

import "fmt"

// Basic function
func add(a int, b int) int {
	return a + b
}

// Shorthand when params share a type
func multiply(a, b int) int {
	return a * b
}

// Multiple return values (very common in Go, e.g. value, error)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

// Named return values (returned automatically with bare `return`)
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // returns min and max
}

// Variadic function: accepts any number of arguments
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Function as a parameter (higher-order function)
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func main() {
	fmt.Println("add:", add(2, 3))
	fmt.Println("multiply:", multiply(4, 5))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("divide:", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	min, max := minMax([]int{4, 1, 9, 2, 7})
	fmt.Println("min:", min, "max:", max)

	fmt.Println("sum:", sum(1, 2, 3, 4, 5))

	square := func(x int) int { return x * x } // anonymous function
	fmt.Println("applyTwice(square, 3):", applyTwice(square, 3))
}
