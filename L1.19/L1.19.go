package main

import "fmt"

func revers(str []rune) []rune {
	i := 0
	j := len(str) - 1
	for i < j {

		str[i], str[j] = str[j], str[i]
		i++
		j--

	}
	return str

}

func main() {

	str := []rune("artem_!@Sinev")
	fmt.Println(str)
	fmt.Println(revers(str))
}
