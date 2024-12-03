package main

import (
	"fmt"
	"math/rand"
	"time"
)

func send(ch chan<- int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(500 * time.Millisecond)
	}
}

func read(ch <-chan int) {
	for {
		val, ok := <-ch
		if !ok {
			return
		}
		fmt.Printf("Reading %d...\n", val)
	}
}

// How many ms it needs for init and run gouroutineS?
func main() {
	n := 5 * time.Second

	ch := make(chan int)

	go send(ch)
	go read(ch)

	time.Sleep(n)
}
