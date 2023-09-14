package thread

import "sync"

type ThreadCounter struct {
	v  int
	mu sync.Mutex
}

func NewThreadCounter() *ThreadCounter {
	return &ThreadCounter{}
}

func (t *ThreadCounter) Inc() {
	t.mu.Lock()
	t.v++
	t.mu.Unlock()
}

func (t *ThreadCounter) Dec() {
	t.mu.Lock()
	t.v--
	t.mu.Unlock()
}

func (t *ThreadCounter) Value() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.v
}
