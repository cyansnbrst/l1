package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) []byte {
	slice := make([]byte, size)

	for i := range slice {
		slice[i] = 'a'
	}

	return slice
}

// func someFunc() {
// 	v := createHugeString(1 << 10)
// //	justString contains a link to all 1 << 10 array
// 	justString = v[:100]
// }

func someFunc() string {
	v := createHugeString(1 << 10)

	justString := make([]byte, 100)
	copy(justString, v[:100])

	return string(justString)
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
	fmt.Println(len(justString))
}
