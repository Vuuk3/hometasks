package main

import (
	"fmt"
	"math"
)

var a, b, c, d float64

func main() {
	fmt.Scan(&a, &b, &c)
	d = math.Pow(b, 2) - 4 * a * c
	if d > 0 {
		fmt.Println((-b - math.Sqrt(d)) / (2 * a), (-b + math.Sqrt(d)) / (2 * a))
	} else if d == 0 {
		fmt.Println(-b / (2 * a))
	} else {
		fmt.Println((complex(-b, -math.Sqrt(-d))) / complex(2 * a, 0), complex(-b, math.Sqrt(-d)) / complex(2 * a, 0))
	}
}