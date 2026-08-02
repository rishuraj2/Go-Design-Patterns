package counter

import (
	"errors"
	"sync"
)

type Counter struct {
	value int
	mu sync.RWMutex
}

var (
	instance *Counter
	once     sync.Once
)

func GetCounter() *Counter {
	once.Do(func() {
		instance = &Counter{
			value: 0,
		}
	})

	return instance
}

func (c *Counter) Increment() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.value == 9999 {
		return errors.New("max limit reached")
	}

	c.value++

	return nil
}

func (c *Counter) GetCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
