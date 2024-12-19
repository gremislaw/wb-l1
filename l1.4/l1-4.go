package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type WorkersPool struct {
	jobs chan int
}

func (w *WorkersPool) Worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case job, ok := <-w.jobs:
				if !ok { // канал закрыт
					return
				}
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Worker %d: %d\n", id, job)
			case <-ctx.Done(): // контекст завершен
				return
			}
		}
	}()
}

func main() {
	var n int
	fmt.Print("Workers' count: ")
	fmt.Scan(&n)

	wp := WorkersPool{
		jobs: make(chan int),
	}

	// Контекст для завершения горутин
	ctx, cancel := context.WithCancel(context.Background())

	// Запись сигналов ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	var wg sync.WaitGroup
	for i := 1; i <= n; i++ {
		wp.Worker(ctx, &wg, i)
	}

	// Генерация данных
	go func() {
		for {
			select {
			case <-sigChan:
				// При получении сигнала о завершении программы
				cancel()
				close(wp.jobs)
				fmt.Println("Shutting down...")
				return
			default:
				// Отправляем данные в канал
				wp.jobs <- rand.Int()
			}
		}
	}()

	wg.Wait() // Ждём завершения всех воркеров
}
