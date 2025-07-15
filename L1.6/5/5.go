package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// через контекст 2 способ
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("контекст отменен, завершаемся")
				return
			default:
				fmt.Println("работает")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
