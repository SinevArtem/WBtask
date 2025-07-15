package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("введите количество секунд: ")
	var N int
	fmt.Scan(&N)

	ch := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case ch <- i:
				fmt.Println("записали в канал ", i)
				i++
				time.Sleep(1 * time.Second)
			case <-ctx.Done():
				fmt.Println("время истекло, закрываем запись")
				close(ch)
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					fmt.Println("канал закрыт")
					return
				}
				fmt.Println(i)

			}
		}

	}()

	wg.Wait()
}
