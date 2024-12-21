// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 4
	b := 11

	fmt.Println(a, b)

	// Множественное присваивание
	a, b = b, a

	fmt.Println(a, b)
}
