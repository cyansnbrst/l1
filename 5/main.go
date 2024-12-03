package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func send(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- rand.Intn(100)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func read(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Reading %d...\n", val)
		}
	}
}

func main() {
	n := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), n)
	defer cancel()

	ch := make(chan int)

	go send(ctx, ch)
	go read(ctx, ch)

	<-ctx.Done()
}
