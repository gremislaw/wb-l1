// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// Функция createHugeString создает строку большого размера,
// и в justString = v[:100] строка v остается в памяти до тех пор,
// пока на нее есть ссылки. Это может привести к ненужному потреблению памяти,
// еще минус. что justString - глобальная переменная, существующая на протяжении работы всей программы.
// Для исправления надо не пользоваться срезом строки и использовать локальные переменные

package main

import (
	"fmt"
	"runtime"
)

// Функция для создания большой строки
func createHugeString(size int) string {
	return string(make([]byte, size)) // Создаем строку заданного размера
}

func someFunc() string {
	v := createHugeString(1 << 10)

	// Выделим новую память для срезанной строки, чтобы не ссылаться на большую строку
	tmp := make([]byte, 100) // Новый срез размером 100
	copy(tmp, v[:100])       // Сделаем реальную аллокацию данных
	return string(tmp)
}

// Функция для проверки использования памяти при вызове функции someFunc, где мы берем 100 символов из строки 1000 символов
func checkMemoryUsage1() {
	fmt.Println(someFunc())

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Используемая память: %v байт\n", memStats.Alloc)
}

// Функция для проверки использования памяти при создании строки 100 символов
func checkMemoryUsage2() {
	fmt.Println(createHugeString(100))

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Используемая память: %v байт\n", memStats.Alloc)
}

func main() {
	checkMemoryUsage1()
	checkMemoryUsage2()
	// Должно быть одинаково плюс минус
}
