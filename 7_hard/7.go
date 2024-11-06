package main

import (
	"fmt"
	"math"
)


var number int


func main() {
	fmt.Scan(&number)
	flag := true
	for i := 2; float64(i) < math.Sqrt(float64(number)) + 1; i++ {
		if number % i == 0 {
			flag = false
			break
		}
	}
	fmt.Println(flag)
}