package main

import (
	"fmt"
	"runtime"
	"sync"
)

// source: https://www.youtube.com/watch?v=qDgs2FJakVQ&list=WL&index=1&ab_channel=ThisIsIT

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * job
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		for _, num := range numbers {
			jobs <- num
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
