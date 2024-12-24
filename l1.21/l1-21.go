// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Target представляет интерфейс, с которым будет работать система
type Target interface {
	Request() string
}

// Adaptee - несовместивая структура к нашей системе
type Adaptee struct {
}

// NewAdapter - это конструктор Adapter
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// Несоответствующий интерфейсу функционал
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter реализует интерфейс target
type Adapter struct {
	*Adaptee
}

// Адаптирование несовместимого функционала
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	target := NewAdapter(adaptee)

	fmt.Println(target.Request())
}
