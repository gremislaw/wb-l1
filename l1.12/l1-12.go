// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

//// ПРИ ЗАПУСКЕ НЕ ЗАБЫВАТЬ УКАЗЫВАТЬ ФАЙЛ "set.go"
//// go run l1.12/l1-12.go l1.12/set.go

package main

func main() {
	set := NewSet("cat", "cat", "dog", "cat", "tree")

	set.Print()
}
