package main

import "fmt"

type Student struct {
}

func (s *Student) Study() {
	fmt.Println("учится")
}

type Target interface {
	Work()
}

type Adapter struct {
	*Student
}

func (a *Adapter) Work() {
	a.Study()
}

func NewAdapter(student *Student) Target {
	return &Adapter{Student: student}
}

func main() {
	adapter := NewAdapter(&Student{})
	adapter.Work()
}
