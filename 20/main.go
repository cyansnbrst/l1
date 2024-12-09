package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// bytes.Split?
	buff := strings.Fields(s)
	n := len(buff)

	for i := 0; i < n/2; i++ {
		buff[i], buff[n-1-i] = buff[n-1-i], buff[i]
	}

	// bytes.Join?
	return strings.Join(buff, " ")

}

func main() {
	myString := "snow dog sun"
	fmt.Println(reverseWords(myString))
}
