package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
}

func main() {

	go func() {
		fmt.Println("запуск")
		sleep(3 * time.Second)
		fmt.Println("дождался")
	}()

	time.Sleep(4 * time.Second)

}
