// Реализовать пересечение двух неупорядоченных множеств.

//// ПРИ ЗАПУСКЕ НЕ ЗАБЫВАТЬ УКАЗЫВАТЬ ФАЙЛ "set.go"
//// go run l1.11/l1-11.go l1.11/set.go

package main

func main() {
	set1 := NewSet(1, 2, 2, 4, 5, 3, 1) // Множество
	set1.Print()

	set2 := NewSet(3, 2, 5, 6, 6, 8) // Множество
	set2.Print()

	res := set1.Intersect(set2) // Пересечение множеств
	res.Print()
}
