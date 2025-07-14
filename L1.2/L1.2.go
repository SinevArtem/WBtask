package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, i := range a {
		wg.Add(1)
		num := i
		go func() {
			defer wg.Done()

			fmt.Println(num * num)
		}()
	}

	wg.Wait()

}
