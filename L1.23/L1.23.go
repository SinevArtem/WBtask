package main

import "fmt"

func del(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr
	}
	copy(arr[index:], arr[index+1:])
	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arr2 := del(arr, 5)
	fmt.Println(arr2)

}
