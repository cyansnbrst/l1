package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, elem := range strings {
		set[elem] = struct{}{}
	}

	for key := range set {
		fmt.Println(key)
	}
}
