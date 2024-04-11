package slidingwindow

import (
	"context"
	"encoding/gob"
	"log"
	"os"
	"sync"
	"time"
)

type Counter struct {
	mu               *sync.Mutex
	Interval         time.Duration
	NumberOfTicks    uint
	RequestBuffer    []uint
	CurrentBufferPos uint
	CallsDuringTick  uint
	TotalCalls       uint
	saveFilePath     string
}

func NewCounter(ctx context.Context, interval time.Duration, numberOfTicks uint, saveFilePath string) *Counter {
	counter := &Counter{
		mu:            &sync.Mutex{},
		Interval:      interval,
		NumberOfTicks: numberOfTicks,
		RequestBuffer: make([]uint, numberOfTicks),
		saveFilePath:  saveFilePath,
	}

	counter.startTicker(ctx)

	return counter
}

func (c *Counter) IncreaseCount() uint {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.CallsDuringTick++
	c.TotalCalls++

	return c.TotalCalls
}

func (c *Counter) GetCount() uint {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.TotalCalls
}

func (c *Counter) tick() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.TotalCalls -= c.RequestBuffer[c.CurrentBufferPos]
	c.RequestBuffer[c.CurrentBufferPos] = c.CallsDuringTick

	c.CallsDuringTick = 0
	c.CurrentBufferPos++
	c.CurrentBufferPos %= uint(len(c.RequestBuffer))
}

func (c *Counter) saveToFile() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	f, err := os.OpenFile(c.saveFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := gob.NewEncoder(f)
	encoder.Encode(c)

	return nil

}

func (c *Counter) readFromFile() error {
	c.mu.Lock()

	temp := c.mu
	filepath := c.saveFilePath
	defer temp.Unlock()

	f, err := os.Open(c.saveFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	decoder := gob.NewDecoder(f)

	tempCounter := &Counter{}
	decoder.Decode(tempCounter)

	if tempCounter.NumberOfTicks != c.NumberOfTicks {
		log.Println("Incompatible save file, continuing with empty data")

		return nil
	}

	*c = *tempCounter
	c.mu = temp
	c.saveFilePath = filepath

	return nil
}

func (c *Counter) startTicker(ctx context.Context) {
	c.readFromFile()

	go func() {
		ticker := time.NewTicker(c.Interval / time.Duration(c.NumberOfTicks))

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				c.tick()
				go c.saveToFile()
			}
		}
	}()

}
