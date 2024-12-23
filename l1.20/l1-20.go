// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	resWords := make([]string, 0)

	// Разбивка строки на слайс из слов
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		resWords = append(resWords, words[i])
	}

	// Объединение слайса слов в строку
	res := strings.Join(resWords, " ")
	return res
}

func main() {
	example := "snow dog sun"
	fmt.Println(reverseWords(example))
}
