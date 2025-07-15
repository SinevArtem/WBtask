package main

import (
	"fmt"
	"sync"
	"time"
)

// через канал уведомлений

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("закрываем горутину")
				return
			default:
				fmt.Println("работаю")
				time.Sleep(time.Second)
			}
		}

	}()

	time.Sleep(7 * time.Second)
	close(done)
	wg.Wait()
}
