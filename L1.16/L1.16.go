package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, right, equal []int

	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num > pivot:
			right = append(right, num)
		default:
			equal = append(equal, num)
		}
	}

	return append(append(quickSort(left), equal...), quickSort(right)...)
}

func main() {
	arr := []int{6, 80, 2, 55, 57, 1, 0, 100, 34}

	fmt.Println(arr)
	fmt.Println(quickSort(arr))
}
