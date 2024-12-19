#!/bin/bash

for i in $(seq 1 26); do
    # Форматируем номер с ведущим нулем
    formatted_number=$(printf "%d" $i)
    
    # Создаем папку с префиксом 'l1.'
    mkdir "l1.$formatted_number"
    
    # Создаем файл внутри каждой папки
    file="l1.$formatted_number/l1-$formatted_number.go"
    touch "$file"

    # Записываем содержимое в файл
    echo 'package main' > "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'import "fmt"' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'func main() {' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '}' >> "$file"

done