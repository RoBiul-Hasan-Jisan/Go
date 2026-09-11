package main

/*
select: like a switch statement but for channel operations. Blocks until
ONE of its cases can proceed; if multiple are ready, one is chosen at random.
Used heavily for: timeouts, cancellation, fan-in/fan-out patterns, and
combining goroutines' results.
*/

import (
	"fmt"
	"time"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	merged := make(chan string)
	go func() {
		defer close(merged)
		open1, open2 := true, true
		// loop while at least one source channel is still open;
		// the condition is re-checked at the TOP of every iteration,
		// so `continue` always lands somewhere safe (unlike checking
		// after the select, which `continue` would skip)
		for open1 || open2 {
			select {
			case v, ok := <-ch1:
				if !ok {
					open1 = false
					ch1 = nil // disable this case: receiving from nil blocks forever
					continue
				}
				merged <- v
			case v, ok := <-ch2:
				if !ok {
					open2 = false
					ch2 = nil
					continue
				}
				merged <- v
			}
		}
	}()
	return merged
}

// A real worker-pool: N workers pulling jobs from a shared channel,
// bounded concurrency, with a timeout guard using select.
func processJob(job int) int {
	time.Sleep(30 * time.Millisecond)
	return job * job
}

func runWorkerPool(numWorkers int, jobs []int) []int {
	jobsCh := make(chan int, len(jobs))
	resultsCh := make(chan int, len(jobs))

	for w := 0; w < numWorkers; w++ {
		go func() {
			for j := range jobsCh {
				resultsCh <- processJob(j)
			}
		}()
	}

	for _, j := range jobs {
		jobsCh <- j
	}
	close(jobsCh)

	results := make([]int, 0, len(jobs))
	for i := 0; i < len(jobs); i++ {
		results = append(results, <-resultsCh)
	}
	return results
}

func main() {
	// ---------- select with timeout ----------
	slowCh := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slowCh <- "finally done"
	}()

	select {
	case res := <-slowCh:
		fmt.Println("Got result:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out waiting for result")
	}

	// ---------- select with default (non-blocking check) ----------
	nonBlocking := make(chan int)
	select {
	case v := <-nonBlocking:
		fmt.Println("Received:", v)
	default:
		fmt.Println("No value ready, moving on")
	}

	// ---------- fan-in ----------
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "from source 1"
		close(c1)
	}()
	go func() {
		c2 <- "from source 2"
		close(c2)
	}()
	for v := range fanIn(c1, c2) {
		fmt.Println("Fan-in received:", v)
	}

	// ---------- worker pool ----------
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	results := runWorkerPool(4, jobs)
	fmt.Println("Worker pool results:", results)
}
