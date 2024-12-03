package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := range numbers {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("%d ", numbers[idx])
		}(i)
	}

	wg.Wait()
}
