package main

/*
Goroutines: lightweight threads managed by the Go runtime, not the OS.
Start one with the `go` keyword. main() will not wait for goroutines
to finish unless you explicitly synchronize (sync.WaitGroup, channels, etc.)
*/

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(label string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(label, i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	// WITHOUT synchronization, main could exit before goroutines finish.
	go printNumbers("goroutine-A")
	go printNumbers("goroutine-B")

	// Give them a moment to run (bad practice - just for a first demo)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("---- now using sync.WaitGroup (the correct way) ----")

	// WaitGroup: proper way to wait for goroutines to finish
	var wg sync.WaitGroup
	results := make([]int, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			results[idx] = idx * idx // each goroutine computes a square
		}(i) // pass i as an argument to avoid the classic loop-variable bug
	}

	wg.Wait() // blocks until all 5 goroutines call wg.Done()
	fmt.Println("Squares computed concurrently:", results)

	// Mutex: protect shared state from concurrent access (race conditions)
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("Final counter (should be 100):", counter)
}
