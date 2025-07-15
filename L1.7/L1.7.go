package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mp := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mu.Lock()
			mp[i] = i * i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	mu.Lock()

	for _, i := range mp {
		fmt.Println(i)
	}

	mu.Unlock()

}
