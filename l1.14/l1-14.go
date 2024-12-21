// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any:
		fmt.Println("channel")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var intValue interface{} = 17
	printType(intValue)

	var stringValue interface{} = "17"
	printType(stringValue)

	var boolValue interface{} = true
	printType(boolValue)

	var chanValue1 interface{} = make(chan any)
	printType(chanValue1)

	var chanValue3 interface{} = 1.7
	printType(chanValue3)
}
