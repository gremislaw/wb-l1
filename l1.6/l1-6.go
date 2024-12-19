package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

/*** Использование канала ***/

func quit1() {
	ch := make(chan struct{})

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("some information")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ch)

	time.Sleep(2 * time.Second)
	close(ch)
	fmt.Println("closing goroutine by closing channel")
	time.Sleep(1 * time.Second)
}

/*** Использование контекста ***/

func quit2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("some information")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("closing goroutine by cancelling context")
	time.Sleep(1 * time.Second)
}

/*** Использование флага ***/

func quit3() {
	var stop atomic.Bool
	stop.Store(false)

	go func(stop *atomic.Bool) {
		for {
			if stop.Load() {
				return
			}
			fmt.Println("some information")
			time.Sleep(500 * time.Millisecond)
		}
	}(&stop)

	time.Sleep(2 * time.Second)
	stop.Store(true)
	fmt.Println("closing goroutine by stop flag")
	time.Sleep(1 * time.Second)
}

/*** Использование таймера ***/

func quit4() {
	q := make(chan struct{})
	timer := time.NewTimer(3 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func(timer *time.Timer, ticker *time.Ticker) {
		for {
			select {
			case <-timer.C:
				close(q)
				return
			case <-ticker.C:
				fmt.Println("some information")
			}
		}
	}(timer, ticker)
	<-q
	fmt.Println("closing goroutine by timer")
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("channel to close goroutine")
	quit1()

	fmt.Println("context to close goroutine")
	quit2()

	fmt.Println("flag to close goroutine")
	quit3()

	fmt.Println("timer to close goroutine")
	quit4()
}
