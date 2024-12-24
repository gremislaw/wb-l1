// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func IsBeauty(s string) bool {
	// Приведение к нижнему регистру для регистронезависимой проверки
	s = strings.ToLower(s)

	// Создадим множество для отслеживания встреченных символов
	checkMap := make(map[rune]struct{}, 0)
	for _, sym := range s {
		if _, ok := checkMap[sym]; !ok {
			checkMap[sym] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsBeauty("abcd"))      // true
	fmt.Println(IsBeauty("abCdefAaf")) // false
	fmt.Println(IsBeauty("aabcd"))     // false
}
