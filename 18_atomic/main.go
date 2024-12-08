package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value int64
}

func (c *counter) add() {
	atomic.AddInt64(&c.value, 1)
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
