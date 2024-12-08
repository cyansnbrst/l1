package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// Cancelling using a channel
func channelCancellation() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("1 - Gouroutine stopped via channel")
				return
			default:
				fmt.Println("1 - Gouroutining...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stop <- true
}

// Cancelling using a cancel context (also could be context with timeout or a deadline)
func contextCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("2 - Gouroutine stopped via context")
				return
			default:
				fmt.Println("2 - Gouroutining...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
}

// Cancelling using an atomic variable
func atomicVariableCancellation() {
	var stop int32 = 0

	go func() {
		for atomic.LoadInt32(&stop) == 0 {
			fmt.Println("3 - Gouroutining...")
			time.Sleep(time.Second)
		}
		fmt.Println("3 - Gouroutine stopped via atomic variable")
	}()

	time.Sleep(3 * time.Second)
	atomic.StoreInt32(&stop, 1)
}

// Cancelling using time.After()
func timeAfterCancellation() {
	done := make(chan bool)
	go func() {
		timer := time.After(3 * time.Second)
		for {
			select {
			case <-timer:
				fmt.Println("4 - Gouroutine stopped via time.After")
				done <- true
				return
			default:
				fmt.Println("4 - Gouroutining...")
				time.Sleep(time.Second)
			}
		}
	}()

	<-done
}

// Cancelling using time.NewTimer
func timeNewTimerCancellation() {
	done := make(chan bool)
	go func() {
		timer := time.NewTimer(3 * time.Second)
		for {
			select {
			case <-timer.C:
				fmt.Println("5 - Gouroutine stopped via time.NewTimer")
				done <- true
				return
			default:
				fmt.Println("5 - Gouroutining...")
				time.Sleep(time.Second)
			}
		}
	}()

	<-done
}

func main() {
	channelCancellation()
	contextCancellation()
	atomicVariableCancellation()
	timeAfterCancellation()
	timeNewTimerCancellation()
}
