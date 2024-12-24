// Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeElement(arr []int, index int) ([]int, error) {
	if index < 0 || index >= len(arr) {
		return arr, fmt.Errorf("index out of range") // Вернем тот же массив
	}
	return append(arr[:index], arr[index+1:]...), nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	var err error
	arr, err = removeElement(arr, 3) // Удаление элемента под индексом 3
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(arr)
}
