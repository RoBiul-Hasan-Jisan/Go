package main

/*
Go has no exceptions (try/catch). Errors are ordinary values, usually
returned as the LAST return value and checked explicitly with `if err != nil`.
*/

import (
	"errors"
	"fmt"
)

// Custom error type: implement the `error` interface -> Error() string
type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

func findUser(id int) (string, error) {
	if id == 1 {
		return "Alice", nil
	}
	return "", &NotFoundError{Resource: fmt.Sprintf("user %d", id)}
}

// Sentinel error (predefined, compared with errors.Is)
var ErrInsufficientFunds = errors.New("insufficient funds")

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, ErrInsufficientFunds
	}
	return balance - amount, nil
}

func main() {
	// Basic error check
	name, err := findUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", name)
	}

	// Custom error + errors.As to inspect the concrete type
	_, err = findUser(5)
	var nfErr *NotFoundError
	if errors.As(err, &nfErr) {
		fmt.Println("Custom error detail -> Resource:", nfErr.Resource)
	}

	// Sentinel error + errors.Is
	_, err = withdraw(100, 500)
	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Println("Transaction failed:", err)
	}

	// Wrapping errors with %w (preserves the original error for errors.Is/As)
	wrapped := fmt.Errorf("processing withdrawal failed: %w", ErrInsufficientFunds)
	fmt.Println("Wrapped error:", wrapped)
	fmt.Println("Is it still ErrInsufficientFunds?", errors.Is(wrapped, ErrInsufficientFunds))

	// panic + recover (for truly exceptional, unrecoverable-by-normal-means situations)
	safeDivide(10, 0)
	fmt.Println("Program continues after recovering from panic")
}

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println(a / b) // division by zero panics
}
