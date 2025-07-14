package main

import (
	"fmt"
	"time"
)

func Worker(numWorker int, ch <-chan int) {
	for i := range ch {

		fmt.Printf("%d читает из канала ресурс: %d \n", numWorker, i)
		time.Sleep(3 * time.Second)
		fmt.Printf("%d закончил работу \n", numWorker)
	}

}

func main() {
	fmt.Print("Введите количестов воркеров: ")
	var numWorkers int

	fmt.Scan(&numWorkers)

	ch := make(chan int)

	for i := 0; i < numWorkers; i++ {

		go Worker(i, ch)
	}

	i := 0
	for {

		ch <- i
		i++
		time.Sleep(1 * time.Second)

	}

}
