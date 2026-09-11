package main

/*
Channels: typed pipes that goroutines use to communicate and synchronize.
ch := make(chan int)       // unbuffered - send blocks until someone receives
ch := make(chan int, 5)    // buffered - send blocks only when buffer is full
"Don't communicate by sharing memory; share memory by communicating." - Go proverb
*/

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // keeps receiving until the channel is closed
		time.Sleep(20 * time.Millisecond)
		results <- j * 2
		_ = id
	}
}

func main() {
	// ---------- Basic unbuffered channel ----------
	messages := make(chan string)
	go func() {
		messages <- "hello from goroutine" // blocks until received
	}()
	msg := <-messages
	fmt.Println("Received:", msg)

	// ---------- Buffered channel ----------
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("Buffered values:", <-buffered, <-buffered, <-buffered)

	// ---------- Closing a channel + range ----------
	nums := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			nums <- i
		}
		close(nums) // signals "no more values"
	}()
	for n := range nums { // loop ends automatically when channel is closed
		fmt.Println("Got:", n)
	}

	// ---------- select: wait on multiple channels ----------
	c1 := make(chan string)
	c2 := make(chan string)
	go func() { time.Sleep(50 * time.Millisecond); c1 <- "from c1" }()
	go func() { time.Sleep(100 * time.Millisecond); c2 <- "from c2" }()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println(m1)
		case m2 := <-c2:
			fmt.Println(m2)
		}
	}

	// ---------- Worker pool pattern ----------
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}
