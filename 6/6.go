package main

import (
	"fmt"
	"math"
)


var a, b float64
var opperation string

func main() {
	fmt.Scan(&a, &b, &opperation)
	if opperation == "+" {
		fmt.Println(a + b)
	}
	if opperation == "-" {
		fmt.Println(a - b)
	}
	if opperation == "*" {
		fmt.Println(a * b)
	}
	if opperation == "/" {
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("На ноль делить нельзя")
		}
	}
	if opperation == "^" {
		fmt.Println(math.Pow(a, b))
	}
	if opperation == "%" {
		if float64(int(a)) == a && float64(int(b)) == b{
			fmt.Println(int(a) % int(b))
		} else {
			fmt.Println("Неверный формат ввода")
		}
	}
}