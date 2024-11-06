package main

import (
	"fmt"
)

var a int64


func main() {
	fmt.Scan(&a)
	fmt.Println(a % 100 == 0 && a % 400 == 0 || a % 4 == 0 && a % 100 != 0)
}