package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var result int64

	var wg sync.WaitGroup

	wg.Add(len(arr))

	for _, el := range arr {
		go func(el int) {
			defer wg.Done()
			atomic.AddInt64(&result, int64(el*el))
		}(el)
	}

	wg.Wait()
	fmt.Println(result)
}
