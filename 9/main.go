package main

import "fmt"

// source : https://kovardin.ru/articles/go/modeli-konkurentnosti-v-go/

func read(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	chan1 := read(arr...)
	chan2 := square(chan1)

	for elem := range chan2 {
		fmt.Println(elem)
	}
}
