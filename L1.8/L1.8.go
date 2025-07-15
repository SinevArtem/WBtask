package main

import (
	"fmt"
	"strconv"
)

func checkBit(num int64, i uint) int {
	if (num & (1 << i)) != 0 {
		return 1
	}
	return 0
}

func main() {
	var num int64
	fmt.Println("введите число ")
	fmt.Scan(&num)

	fmt.Println(strconv.FormatInt(num, 2))
	var i uint

	fmt.Println("введите номер бита ")
	fmt.Scan(&i)

	if checkBit(num, i) == 1 {
		num = num &^ (1 << i)
	} else {
		num = num | (1 << i)
	}
	fmt.Println(num)
	fmt.Println(strconv.FormatInt(num, 2))

}
