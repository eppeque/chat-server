package utils

import "sync"

type ThreadCounter struct {
	counter int
	mu      sync.Mutex
}

func NewThreadCounter() *ThreadCounter {
	return &ThreadCounter{}
}

func (c *ThreadCounter) Increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *ThreadCounter) Decrement() {
	c.mu.Lock()
	c.counter--
	c.mu.Unlock()
}

func (c *ThreadCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}
