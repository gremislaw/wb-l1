// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

// Установка 1 в i-й бит
func SetOneBit(n *int64, i int) {
	*n |= (1 << (i - 1))
}

// Установка 0 в i-й бит
func SetZeroBit(n *int64, i int) {
	*n &= ^(1 << (i - 1))
}

// Получение i-й бита
func GetBit(n int64, i int) int64 {
	return (n >> (i - 1)) & 1
}

func main() {
	var n int64 = 10

	i := 3

	fmt.Printf("\nbinary: %.8b, decimal: %d\n", n, n)
	fmt.Printf("Value of %d bit: %d\n\n", i, GetBit(n, i))

	SetOneBit(&n, i)
	fmt.Printf("Set 1 to the %d bit\nbinary: %.8b, decimal: %d\n", i, n, n)
	fmt.Printf("Value of %d bit: %d\n\n", i, GetBit(n, i))

	SetZeroBit(&n, i)
	fmt.Printf("Set 0 to the %d bit\nbinary: %.8b, decimal: %d\n", i, n, n)
	fmt.Printf("Value of %d bit: %d\n\n", i, GetBit(n, i))
}
