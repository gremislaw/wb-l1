package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHello() {
	fmt.Printf("Hello! My name is %v and i'm %v years old\n", h.name, h.age)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		name: "Dazhy",
		age:  20,
	}

	action := Action{
		Human: human,
	}

	action.SayHello()
}
