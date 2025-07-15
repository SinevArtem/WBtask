package main

import (
	"fmt"
	"sort"
)

func main() {
	res := make(map[int][]float64)
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sort.Float64s(arr)

	for i, value := range arr {
		res[int(arr[i]/10)*10] = append(res[int(arr[i]/10)*10], value)
	}

	fmt.Println(res)
}
