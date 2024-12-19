package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var duration int
	fmt.Print("Timer: ")
	fmt.Scan(&duration)

	// Таймер на завершение программы
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	ch := make(chan int)

	// Горутина для отправки данных в канал
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for {
			fmt.Printf("Got: %d\n", <-ch)
		}
	}()

	<-timer.C
	fmt.Println("Timeout. Shutting down...")
}
