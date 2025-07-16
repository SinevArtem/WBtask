package main

import "fmt"

func revers(str []rune, start, end int) []rune {
	for start < end {

		str[start], str[end] = str[end], str[start]
		start++
		end--

	}
	return str

}

func main() {
	str := []rune("snow dog sun")
	l := len(str)
	str = revers(str, 0, l-1)

	start := 0
	for i := 0; i <= l; i++ {
		if i == l || str[i] == 32 {
			revers(str, start, i-1)
			start = i + 1
		}
	}

	fmt.Println(string(str))
}
