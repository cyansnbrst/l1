package main

import "fmt"

func removeElem(s []int, idx int) []int {
	if idx < 0 || idx >= len(s) {
		fmt.Println("index out of range")
		return s
	}

	return append(s[:idx], s[idx+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = removeElem(slice, 4)
	fmt.Println(slice)
}
