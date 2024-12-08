package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int64
	mu    sync.Mutex
}

func (c *counter) add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	var wg sync.WaitGroup
	c := counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.add()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("counter: %d", c.value)
}
