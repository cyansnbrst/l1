package main

import "fmt"

func main() {
	a := int64(7)
	b := int64(10)

	fmt.Printf("a, b: %d, %d\n", a, b)

	b = a ^ b
	a = a ^ b
	b = a ^ b

	fmt.Printf("a, b: %d, %d\n", a, b)
}
