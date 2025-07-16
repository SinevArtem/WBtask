package main

import (
	"fmt"
	"unicode"
)

func f(s string) bool {
	mp := make(map[rune]bool)

	for _, char := range s {
		low := unicode.ToLower(char)
		if mp[low] {
			return false
		}
		mp[low] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(f(str1))
	fmt.Println(f(str2))
	fmt.Println(f(str3))
}
