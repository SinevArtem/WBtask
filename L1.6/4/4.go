package main

import (
	"fmt"
	"runtime"
	"sync"
)

// runtime.Goexit()

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("полнит все defer")
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		runtime.Goexit()

		fmt.Println("меня не выполнит") // не выполнит, прервят выполнение горутины

	}()

	wg.Wait()
}
