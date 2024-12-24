// Реализовать пересечение двух неупорядоченных множеств.

package main

import "l1/l1.11/set"

func main() {
	set1 := set.NewSet(1, 2, 2, 4, 5, 3, 1) // Множество
	set1.Print()

	set2 := set.NewSet(3, 2, 5, 6, 6, 8) // Множество
	set2.Print()

	res := set1.Intersect(set2) // Пересечение множеств
	res.Print()
}
