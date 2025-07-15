package main

import "fmt"

func main() {
	whatIsIt := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case chan interface{}:
			fmt.Println("канал")
		default:
			fmt.Println("не знаю что это", v)
		}
	}

	whatIsIt(1)
	whatIsIt("f")
	whatIsIt(true)
	whatIsIt(make(chan interface{}))
	whatIsIt(1.4)
}
