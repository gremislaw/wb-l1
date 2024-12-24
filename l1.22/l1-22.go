// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a и b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := big.NewFloat(0).SetString("9223372036854775808")   // число больше int64
	b, _ := big.NewFloat(0).SetString("100000000000000000001") // число больше 10 ^ 20

	var tmp big.Float
	fmt.Println("a + b = ", tmp.Add(a, b))
	fmt.Println("a - b = ", tmp.Add(a, tmp.Neg(b)))
	fmt.Println("b - a = ", tmp.Add(tmp.Neg(a), b))
	fmt.Println("a * b = ", tmp.Mul(a, b))
	fmt.Println("a / b = ", tmp.Quo(a, b))
	fmt.Println("b / a = ", tmp.Quo(b, a))
}
