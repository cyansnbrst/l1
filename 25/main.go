package main

import (
	"fmt"
	"time"
)

// func sleep(duration time.Duration) {
// 	<-time.After(duration)
// }

func sleep(duration time.Duration) {
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
	}
}

func main() {
	fmt.Println("sleeping...")
	sleep(3 * time.Second)
	fmt.Println("awake!")
}
