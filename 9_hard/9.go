package main

import (
	"fmt"
)


var start, finish uint


func main() {
	fmt.Println("Введите диапазон для вывода таблицы через пробел. (Пример: 5 8)")
	fmt.Scan(&start, &finish)
	start = min(max(start, 1), 10)
	finish = min(finish, 10)
	for i := start; i < finish + 1; i++ {
		for j := start; j < finish + 1; j++ {
			fmt.Print(i, " * ", j, " = ", i * j, " ")
		}
		fmt.Print("\n")
	}
}