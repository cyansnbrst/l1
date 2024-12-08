package main

import "fmt"

func main() {
	set1 := []int64{5, 3, -10, 8, 6}
	set2 := []int64{8, 1, -10, 4, 2}

	m := make(map[int64]int8)

	for _, elem := range set1 {
		m[elem]++
	}

	for _, elem := range set2 {
		m[elem]--
	}

	for key, value := range m {
		if value == 0 {
			fmt.Println(key)
		}
	}
}
