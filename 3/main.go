package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var mu sync.Mutex

	sum := 0

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			square := n * n
			mu.Lock()
			sum += square
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}
