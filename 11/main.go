package main

import "fmt"

func main() {
	set1 := []int64{5, 3, -10, 8, 6, 8}
	set2 := []int64{8, 1, -10, 4, 2}

	m := make(map[int64]int)

	for _, elem := range set1 {
		m[elem]++
	}

	intersection := []int64{}
	for _, elem := range set2 {
		if count, exists := m[elem]; exists && count > 0 {
			intersection = append(intersection, elem)
			m[elem]--
		}
	}

	fmt.Println("Intersection:", intersection)
}
