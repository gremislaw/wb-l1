package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	results := make([]int, len(arr))

	var wg sync.WaitGroup

	wg.Add(len(arr))

	for i, el := range arr {
		go func(i, el int) {
			defer wg.Done()
			results[i] = el * el // мьютекс не нужен т.к. i уникален
		}(i, el)
	}

	wg.Wait()
	fmt.Println(results)
}
