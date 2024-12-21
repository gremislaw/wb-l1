package main

import "fmt"

// Структура множество
type Set struct {
	data map[interface{}]struct{}
}

// Новое множество
func NewSet(arr ...interface{}) *Set {
	data := make(map[interface{}]struct{})
	for _, el := range arr {
		data[el] = struct{}{}
	}
	return &Set{data: data}
}

func (s *Set) GetData() map[interface{}]struct{} {
	return s.data
}

// Проверка элемента во множестве
func (s *Set) Find(n interface{}) bool {
	_, ok := s.data[n]
	return ok
}

// Добавление элемента во множество
func (s *Set) Add(n interface{}) {
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
		fmt.Printf("%v ", el)
	}
	fmt.Print("}\n")
}
