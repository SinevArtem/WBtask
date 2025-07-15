package main

import (
	"fmt"
)

func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		for _, i := range arr {
			ch1 <- i
		}
		close(ch1)

	}()

	go func() {

		for i := range ch1 {
			ch2 <- i * 2
		}
		close(ch2)

	}()

	for i := range ch2 {
		fmt.Println(i)
	}

}
