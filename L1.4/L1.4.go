package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(numWorker int, ch <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("канал закрыт ")
				return
			}
			fmt.Printf("%d читает из канала ресурс: %d \n", numWorker, i)
			time.Sleep(3 * time.Second)
			fmt.Printf("%d закончил работу \n", numWorker)
		case <-ctx.Done():
			fmt.Printf("%d больше не будет работать, контекст закрыт \n", numWorker)
			return
		}

	}

}

func main() {
	fmt.Print("Введите количестов воркеров: ")
	var numWorkers int
	fmt.Scan(&numWorkers)

	var wg sync.WaitGroup

	ch := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, ch, ctx, &wg)
	}

	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
				time.Sleep(1 * time.Second)
			case <-ctx.Done():
				close(ch)
				return
			}

		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	<-done
	fmt.Println("нажал ctrl+c")
	cancel()

	wg.Wait()

}
