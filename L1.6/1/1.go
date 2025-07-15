package main

import (
	"fmt"
	"sync"
)

// по условию

func main() {

	var wg sync.WaitGroup
	i := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			i++
			if i == 100 {
				fmt.Println("завершаем")
				return
			}

		}
	}()

	wg.Wait()
	fmt.Println("закончили работу")

}
