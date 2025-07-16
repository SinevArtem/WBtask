package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	b, _ := new(big.Int).SetString("99999999999999999999999999999999999999999999999999", 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Div(a, b)

	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(prod)
	fmt.Println(quot)
}
