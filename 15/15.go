package main

import (
	"fmt"
)


var number1, number2 int


func main() {
	fmt.Scan(&number1, &number2)
	number1, number2 = max(number1, number2), min(number1, number2)
	for i := 0; number1 % number2 != 0; i++ {
		number1, number2 = number2, number1 % number2
	}
	fmt.Println(number2)
}