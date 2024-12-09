package main

import (
	"fmt"
	"math/big"
)

// Sum numbers
func sum(a, b *big.Int) {
	result := new(big.Int).Add(a, b)
	fmt.Printf("sum: %v + %v = %v\n", a, b, result)
}

// Sub numbers
func sub(a, b *big.Int) {
	result := new(big.Int).Sub(a, b)
	fmt.Printf("sub: %v - %v = %v\n", a, b, result)
}

// Multiply numbers
func multiply(a, b *big.Int) {
	result := new(big.Int).Mul(a, b)
	fmt.Printf("multiplication: %v * %v = %v\n", a, b, result)
}

// Divide numbers
func divide(a, b *big.Int) {
	if b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("zero division")
		return
	}
	aFloat := new(big.Float).SetInt(a)
	bFloat := new(big.Float).SetInt(b)
	result := new(big.Float).Quo(aFloat, bFloat)
	fmt.Printf("division: %v / %v = %v\n", a, b, result)
}

func main() {
	a := big.NewInt(1 << 21)
	b := big.NewInt(1 << 22)

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

	sum(a, b)
	sub(a, b)
	multiply(a, b)
	divide(a, b)
}
