package main

import (
	"fmt"
	"unicode"
)

func uniqueSymbCheck(s string) bool {
	symbols := make(map[rune]struct{})

	for _, r := range s {
		ri := unicode.ToLower(r)
		if _, exists := symbols[ri]; exists {
			return false
		}
		symbols[ri] = struct{}{}
	}
	return true
}

func main() {
	string1 := "abcd"
	string2 := "abcdA"
	string3 := "abCdefaf"

	fmt.Println(uniqueSymbCheck(string1))
	fmt.Println(uniqueSymbCheck(string2))
	fmt.Println(uniqueSymbCheck(string3))
}
