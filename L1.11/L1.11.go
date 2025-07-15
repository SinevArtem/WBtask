package main

import "fmt"

func f(a, b []int) []int {
	set := make(map[int]bool)
	var result []int

	for _, num := range a {
		set[num] = true
	}

	for _, num := range b {
		if set[num] {
			result = append(result, num)
			set[num] = false
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 8, 23, 5, 6, 6, 0}
	b := []int{2, 3, 4, 6, 100, 15, 8}

	fmt.Println(f(a, b))
}
