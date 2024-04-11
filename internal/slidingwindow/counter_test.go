package slidingwindow

import (
	"sync"
	"testing"
	"time"
)

func TestCounterTick(t *testing.T) {

	t.Run("reset value after tick", func(t *testing.T) {
		c := &Counter{
			mu:            &sync.Mutex{},
			Interval:      1 * time.Second,
			NumberOfTicks: 1,
			RequestBuffer: make([]uint, 1),
		}

		c.IncreaseCount()
		c.IncreaseCount()
		value := c.IncreaseCount()
		c.tick()
		c.tick()

		valueAfterTick := c.GetCount()

		if value == valueAfterTick {
			t.Error("Counter didnt change after tick")
		}
	})

	t.Run("keep value after tick", func(t *testing.T) {
		c := &Counter{
			mu:            &sync.Mutex{},
			Interval:      1 * time.Second,
			NumberOfTicks: 2,
			RequestBuffer: make([]uint, 2),
		}

		c.IncreaseCount()
		c.IncreaseCount()
		value := c.IncreaseCount()
		c.tick()
		c.tick()

		valueAfterTick := c.GetCount()

		if value != valueAfterTick {
			t.Errorf("Counter value is not expected, want %d, got %d", value, valueAfterTick)
		}
	})
}
