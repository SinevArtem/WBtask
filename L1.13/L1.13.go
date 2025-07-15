package main

import "fmt"

func main() {

	a := 15
	b := 30

	// b = a + b
	// a = b - a
	// b = b - a

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a, b)
}
