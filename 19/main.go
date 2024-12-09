package main

import (
	"fmt"
)

func reverseString(s string) string {
	buff := []rune(s)
	n := len(buff)

	var temp rune
	for i := 0; i < n/2; i++ {
		temp = buff[i]
		buff[i] = buff[n-1-i]
		buff[n-1-i] = temp
	}

	return string(buff)
}

func main() {
	myString := "главорыба"
	myString2 := "汉语"

	fmt.Println(reverseString(myString))
	fmt.Println(reverseString(myString2))
}
