// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "l1/l1.12/set"

func main() {
	set := set.NewSet("cat", "cat", "dog", "cat", "tree")

	set.Print()
}
