package main

import (
	"fmt"
	"sync"
)

type Count struct {
	i  int
	mu sync.Mutex
}

func main() {

	var wg sync.WaitGroup

	i := Count{i: 0}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 1000000; j++ {
			i.mu.Lock()
			i.i++
			i.mu.Unlock()
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 1000000; j++ {
			i.mu.Lock()
			i.i++
			i.mu.Unlock()
		}

	}()

	wg.Wait()

	fmt.Println(i.i)

}
