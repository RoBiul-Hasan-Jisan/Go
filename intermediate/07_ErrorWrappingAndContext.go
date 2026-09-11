package main

/*
Two "pro" habits that separate intermediate Go devs from beginners:
1. Wrapping errors with context (fmt.Errorf + %w) so you can trace WHERE
   an error came from without losing the original error.
2. Using context.Context to carry cancellation signals & deadlines through
   call chains (very common in servers, HTTP clients, DB calls).
*/

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var ErrRecordNotFound = errors.New("record not found")

func fetchFromDB(id int) error {
	if id != 1 {
		return ErrRecordNotFound
	}
	return nil
}

func getUser(id int) error {
	if err := fetchFromDB(id); err != nil {
		// wrap with context; %w keeps the chain intact for errors.Is/As
		return fmt.Errorf("getUser(%d) failed: %w", id, err)
	}
	return nil
}

func longRunningTask(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second): // simulate slow work
		fmt.Println("task finished normally")
		return nil
	case <-ctx.Done(): // cancelled or timed out
		return ctx.Err()
	}
}

func main() {
	// ---------- Error wrapping ----------
	err := getUser(2)
	fmt.Println("Error:", err)
	fmt.Println("Is it ErrRecordNotFound underneath?", errors.Is(err, ErrRecordNotFound))
	fmt.Println("Unwrapped:", errors.Unwrap(err))

	// ---------- context.WithTimeout ----------
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := longRunningTask(ctx); err != nil {
		fmt.Println("Task error (expected timeout):", err)
	}

	// ---------- context.WithCancel ----------
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("cancelling manually...")
		cancel2()
	}()
	<-ctx2.Done()
	fmt.Println("ctx2 cancelled, reason:", ctx2.Err())

	// ---------- context.WithValue (use sparingly - request-scoped data only) ----------
	// NOTE: in real code, use a custom unexported type as the key (not a
	// plain string) to avoid collisions with other packages. Using a string
	// key here is just to keep the demo short.
	type ctxKey string
	const requestIDKey ctxKey = "requestID"
	ctx3 := context.WithValue(context.Background(), requestIDKey, "abc-123")
	fmt.Println("requestID from context:", ctx3.Value(requestIDKey))
}
