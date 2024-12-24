package main

import (
	"fmt"
	"time"
)

func Sleep(seconds time.Duration) {
	start := time.Now()
	for time.Since(start) < seconds*time.Second {
	}
}

func main() {
	fmt.Println("Запускаем Sleep")
	Sleep(4)
	fmt.Println("Sleep закончился")
}
