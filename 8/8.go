package main

import (
	"fmt"
)

var a1, a2, b1, b2, c1, c2 float64


func main() {
	fmt.Scan(&a1, &a2, &b1, &b2, &c1, &c2)
	fmt.Println(a2 > b1 && a2 > c1 && b2 > a1 && b2 > c1 && c2 > a1 && c2 > b2)
}