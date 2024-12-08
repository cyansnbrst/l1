package main

import "fmt"

func binSearch(array []int, value int) int {
	l := 0
	r := len(array) - 1

	for l <= r {
		m := (l + r) / 2
		if array[m] < value {
			l = m + 1
		} else if array[m] > value {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

func main() {
	arr := []int{5, 14, 26, 36, 48, 51, 62}
	target := 14

	fmt.Println("index:", binSearch(arr, target))
}
