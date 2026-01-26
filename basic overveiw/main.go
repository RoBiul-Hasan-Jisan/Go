/*
for initialization; condition; update {
    // code
}

*/

package main

import "fmt"

func main() {

	// Basic for loop (like 1 to 5)
	fmt.Println("Basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("\n-----------------")

	// For loop like while
	fmt.Println("While style loop:")
	j := 1
	for j <= 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println("\n-----------------")

	//  User input loop (table)
	var num int
	fmt.Print("Enter a number to print table: ")
	fmt.Scan(&num)

	fmt.Println("Table of", num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
	fmt.Println("-----------------")

	// Nested loop
	fmt.Println("Nested loop 3x3:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print("(", i, ",", j, ") ")
		}
		fmt.Println()
	}
	fmt.Println("-----------------")

	// break and continue
	fmt.Println("Break and continue example:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Breaking at i =", i)
			break
		}
		if i%2 == 0 {
			fmt.Println("Skipping even i =", i)
			continue
		}
		fmt.Println("i =", i)
	}
	fmt.Println("-----------------")

	//  For range loop (array)
	fmt.Println("For range loop:")
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("-----------------")

	//  Even/Odd loop
	fmt.Println("Even/Odd check from 1 to 10:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}
	
}
