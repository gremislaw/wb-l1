// Реализовать бинарный поиск встроенными методами языка.

//// ПРИ ЗАПУСКЕ НЕ ЗАБЫВАТЬ УКАЗЫВАТЬ ФАЙЛ "quick_sort.go"
//// go run l1.17/l1-17.go l1.17/quick_sort.go

package main

import "fmt"

func binary_search(arr []int, n int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {

		mid := (left + right) / 2
		if n == arr[mid] {
			return mid
		}

		if n < arr[mid] {
			right = mid - 1
		}

		if n > arr[mid] {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{3, 17, 15, 2, 18, -5, 0, 5, -1, 3, 2, 17, 1, -2}
	fmt.Println(arr)

	quickSort(arr) // Сортировка массива
	fmt.Println(arr)

	n := 3
	fmt.Printf("Elem %d index is %d\n", n, binary_search(arr, n))

	n = 777
	fmt.Printf("Elem %d index is %d\n", n, binary_search(arr, n))

	n = -1
	fmt.Printf("Elem %d index is %d\n", n, binary_search(arr, n))
}
