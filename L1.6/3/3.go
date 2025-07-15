package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// через контекст 1 способ
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("завершаем работу горутины")
		}
	}()

	wg.Wait()

	fmt.Println("завершаем работу")
}
