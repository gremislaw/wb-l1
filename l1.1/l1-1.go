// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

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
