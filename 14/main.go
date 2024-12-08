package main

import "fmt"

func determineType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("%d: int\n", v)
	case string:
		fmt.Printf("%s: string\n", v)
	case bool:
		fmt.Printf("%v: bool\n", v)
	case chan int:
		fmt.Printf("%v: channel\n", v)
	default:
		fmt.Printf("%v: other\n", v)
	}
}

func main() {
	variables := []interface{}{
		42,
		"Hello, World!",
		true,
		make(chan int),
		3.14,
	}

	for _, v := range variables {
		determineType(v)
	}
}
