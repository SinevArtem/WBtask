package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Method() {
	fmt.Println("какой то метод")
}

type Action struct {
	Human
	str string
}

func main() {
	a := Action{
		Human: Human{
			name: "artem",
			age:  20,
		},
		str: "run",
	}

	fmt.Println(a)

	a.Method()

}
