// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	subArr := arr[len(arr)/2]
	l := 0
	r := len(arr) - 1
	for l <= r {
		for arr[l] < subArr {
			l++
		}

		for arr[r] > subArr {
			r--
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	quickSort(arr[l:]) // Cортировка правой половины

	quickSort(arr[:r+1]) // Cортировка левой половины
}

func main() {
	nums1 := []int{3, 17, 15, 2, 18, 0, 5, 3, 3, 2, 17, 1}
	fmt.Print(nums1, " -> my_sort ->\t")

	// Своя реализация
	quickSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{3, 17, 15, 2, 18, 0, 5, 3, 3, 2, 17, 1}
	fmt.Print(nums2, " -> golang_sort ->\t")

	// Пакет sort использует быструю сортировку, но для небольших массивов сортировку вставками
	sort.Ints(nums2)
	fmt.Println(nums2)
}
