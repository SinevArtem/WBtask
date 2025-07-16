package main

import "fmt"

func Search(arr []int, value int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			left = mid + 1
		} else if arr[mid] > value {
			right = mid - 1
		}
	}

	return -1

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(Search(arr, 2))
	fmt.Println(Search(arr, 0))
	fmt.Println(Search(arr, 7))
	fmt.Println(Search(arr, 10))
}
