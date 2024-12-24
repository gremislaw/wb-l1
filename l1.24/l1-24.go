// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

package main

import (
	"fmt"
	"l1/l1.24/point"
)

func main() {
	a := point.NewPoint(5, 3)
	b := point.NewPoint(4, 10)
	
	fmt.Println(a.GetDistance(b))
}
