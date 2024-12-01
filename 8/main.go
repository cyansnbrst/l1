package main

import "fmt"

/*
	source: https://metanit.com/go/tutorial/2.7.php
*/

func setBitOne(number int64, i uint) int64 {
	return number | (1 << i)
}

func setBitZero(number int64, i uint) int64 {
	return number &^ (1 << i)
}

func main() {
	var number int64 = 13
	fmt.Printf("Init number: %d (%b)\n", number, number)
	var bit uint = 0

	toOne := setBitOne(number, bit)
	fmt.Printf("Set %d bit to one: %d (%b)\n", bit, toOne, toOne)

	toZero := setBitZero(number, bit)
	fmt.Printf("Set %d bit to zero: %d (%b)\n", bit, toZero, toZero)
}
