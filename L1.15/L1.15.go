package main

import (
	"fmt"
	"strings"
)

func createHugeString(i int) string {
	return strings.Repeat("a", i)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) < 100 {
		panic("строка слишком короткая")
	}
	return string([]byte(v[:100])) //создает новый массив байтов и новую строку

}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
