package main

/*
defer  -> schedules a call to run right before the surrounding function returns
          (runs in LIFO order - last deferred, first executed)
panic  -> stops normal execution, starts unwinding the stack, running deferred calls
recover -> stops a panic mid-unwind, must be called inside a deferred function
*/

import "fmt"

func cleanupDemo() {
	fmt.Println("1. Entering cleanupDemo")
	defer fmt.Println("4. Deferred: closing resource A")
	defer fmt.Println("3. Deferred: closing resource B")
	fmt.Println("2. Doing work...")
	// deferred calls run in reverse order: B then A -> "3" then "4"
}

// defer is commonly used to guarantee cleanup (files, locks, connections)
func readFile(name string) {
	fmt.Println("Opening", name)
	defer fmt.Println("Closing", name) // guaranteed to run even if code below panics/returns early

	fmt.Println("Reading from", name)
}

// Deferred functions can modify named return values
func doubleWithLog(n int) (result int) {
	defer func() {
		fmt.Println("about to return, doubling result from", result, "to", result*2)
		result = result * 2
	}()
	result = n
	return
}

func riskyOperation(x int) (out int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	if x == 0 {
		panic("x cannot be zero")
	}
	return 100 / x, nil
}

func main() {
	cleanupDemo()
	fmt.Println("-----------------")

	readFile("data.txt")
	fmt.Println("-----------------")

	fmt.Println("doubleWithLog(5):", doubleWithLog(5))
	fmt.Println("-----------------")

	out, err := riskyOperation(0)
	if err != nil {
		fmt.Println("Handled error:", err)
	} else {
		fmt.Println("Result:", out)
	}

	// Program keeps running normally after a recovered panic
	out, err = riskyOperation(5)
	fmt.Println("Result:", out, "Err:", err)
}
