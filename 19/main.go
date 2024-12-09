package main

import (
	"fmt"
)

func reverseString(s string) string {
	buff := []rune(s)
	n := len(buff)

	for i := 0; i < n/2; i++ {
		buff[i], buff[n-1-i] = buff[n-1-i], buff[i]
	}

	return string(buff)
}

func main() {
	myString := "главорыба"
	myString2 := "汉语"

	fmt.Println(reverseString(myString))
	fmt.Println(reverseString(myString2))
}
