package main

/*
sync package essentials beyond basic Mutex/WaitGroup:
- sync.RWMutex: many readers OR one writer (better for read-heavy workloads)
- sync.Once:    run initialization code exactly once, no matter how many goroutines call it
- sync.Map:     a concurrency-safe map (use when you have heavy concurrent read/write)
- atomic:       lock-free primitives for simple counters
*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewSafeCache() *SafeCache {
	return &SafeCache{data: make(map[string]string)}
}

func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.RLock() // multiple readers can hold this simultaneously
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *SafeCache) Set(key, value string) {
	c.mu.Lock() // exclusive - blocks all readers and writers
	defer c.mu.Unlock()
	c.data[key] = value
}

var once sync.Once

func expensiveInit() {
	fmt.Println("Running expensive initialization...")
}

func main() {
	// ---------- RWMutex ----------
	cache := NewSafeCache()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key%d", n), fmt.Sprintf("value%d", n))
		}(i)
	}
	wg.Wait()

	v, _ := cache.Get("key3")
	fmt.Println("Cached value for key3:", v)

	// ---------- sync.Once ----------
	for i := 0; i < 3; i++ {
		once.Do(expensiveInit) // only prints once, regardless of call count
	}

	// ---------- sync.Map ----------
	var sm sync.Map
	sm.Store("a", 1)
	sm.Store("b", 2)
	if val, ok := sm.Load("a"); ok {
		fmt.Println("sync.Map value for 'a':", val)
	}

	// ---------- atomic (lock-free counter) ----------
	var counter int64
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg2.Wait()
	fmt.Println("Atomic counter (should be 1000):", atomic.LoadInt64(&counter))
}
