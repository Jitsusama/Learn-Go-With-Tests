package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increments the counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, &counter, 3)
	})

	t.Run("increments the counter concurrently", func(t *testing.T) {
		var waiter sync.WaitGroup
		count := 1000
		counter := Counter{}

		waiter.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				counter.Inc()
				waiter.Done()
			}()
		}
		waiter.Wait()

		assertCounter(t, &counter, count)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
