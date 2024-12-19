// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	safeMap := new(sync.Map)
	wg := new(sync.WaitGroup)

	// Конкурентная запись в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Store(i, i*i)
			fmt.Printf("wrote %v at %v\n", i*i, i)
		}(i)
	}

	wg.Wait()

	// Чтение из map
	safeMap.Range(func(key, value interface{}) bool {
		fmt.Printf("got %v at %v\n", value, key)
		return true
	})
}
