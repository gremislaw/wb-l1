// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	numChan := make(chan int)
	resChan := make(chan int)

	// Передача числа в первый канал
	go func() {
		for _, x := range arr {
			numChan <- x
		}
		close(numChan)
	}()

	// Передача числа из первого канала во второй
	go func() {
		for x := range numChan {
			resChan <- x * 2
		}
		close(resChan)
	}()

	for x := range resChan {
		fmt.Println(x) // Вывод данных из второго канала
	}
}
