// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func reverse(s string) string {
	res := make([]rune, 0)

	// Преобразование строки в срез рун для поддержки Unicode
	sRune := []rune(s)
	for i := len(sRune) - 1; i >= 0; i-- {
		res = append(res, sRune[i])
	}

	return string(res)
}

func main() {
	example := "главрыба"
	fmt.Println(reverse(example))
}
