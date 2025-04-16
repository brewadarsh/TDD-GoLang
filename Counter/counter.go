package counter

import "sync"

// A counter which is safe to use concurrently.
type Counter struct {
	count int
	mutex sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

// Increment the counter.
func (counter *Counter) Inc() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.count += 1
}

// Get the count value.
func (counter *Counter) Value() int {
	return counter.count
}
