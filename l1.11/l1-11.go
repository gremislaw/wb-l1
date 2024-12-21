// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

// Структура множество
type Set struct {
	data map[int]struct{}
}

// Новое множество
func NewSet(arr ...int) *Set {
	data := make(map[int]struct{})
	for _, el := range arr {
		data[el] = struct{}{}
	}
	return &Set{data: data}
}

func (s *Set) GetData() map[int]struct{} {
	return s.data
}

// Проверка элемента во множестве
func (s *Set) Find(n int) bool {
	_, ok := s.data[n]
	return ok
}

// Добавление элемента во множество
func (s *Set) Add(n int) {
	s.data[n] = struct{}{}
}

// Пересечение множеств
func (s1 *Set) Intersect(s2 *Set) *Set {
	res := NewSet()
	for el := range s1.GetData() {
		res.Add(el)
	}

	for el := range s2.GetData() {
		res.Add(el)
	}

	return res
}

func (s *Set) Print() {
	fmt.Print("{ ")
	for el := range s.data {
		fmt.Printf("%d ", el)
	}
	fmt.Print("}\n")
}

func main() {
	set1 := NewSet(1, 2, 2, 4, 5, 3, 1) // Множество
	set1.Print()

	set2 := NewSet(3, 2, 5, 6, 6, 8) // Множество
	set2.Print()

	res := set1.Intersect(set2) // Пересечение множеств
	res.Print()
}
