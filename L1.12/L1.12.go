package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, i := range str {
		set[i] = struct{}{}

	}

	for i := range set {
		fmt.Println(i)
	}

}
