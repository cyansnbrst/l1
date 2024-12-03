package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// source : https://www.youtube.com/watch?v=qDgs2FJakVQ&list=WL&index=1&ab_channel=ThisIsIT

func worker(id int, in <-chan int, wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-done:
			fmt.Printf("Worker %d endind its work\n", id)
			return
		case num, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processed: %d\n", id, num)
		}
	}
}

func main() {
	var workerCount int
	fmt.Print("Enter workers amount: ")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		fmt.Printf("Incorrect workers amount: %s", err.Error())
		return
	}
	fmt.Printf("Using %d workers\n\n", workerCount)

	input := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, input, &wg, done)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Enter numbers to procces. Ctrl+C for end input")
	var num int
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Scan(&num)
				input <- num
				time.Sleep(1 * time.Second)
			}
		}
	}()

	<-sigChan

	close(done)
	close(input)

	wg.Wait()
}
