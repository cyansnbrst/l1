package main

import "fmt"

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivotIdx := len(array) / 2
	pivot := array[pivotIdx]

	var left, middle, right []int

	for _, num := range array {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			middle = append(middle, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	result := append(quickSort(left), middle...)
	result = append(result, quickSort(right)...)

	return result
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("sorted array:", quickSort(arr))
}
