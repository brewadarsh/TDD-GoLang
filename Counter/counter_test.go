package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment counter 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("Concurrent counter", func(t *testing.T) {
		wanted := 1000
		counter := NewCounter()

		var waitGroup sync.WaitGroup
		waitGroup.Add(wanted)

		for range wanted {
			go func() {
				counter.Inc()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()
		assertCounter(t, counter, wanted)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
