package main

import (
	"fmt"
)


var a, reverse_a string

func main() {
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		reverse_a += string(a[len(a) - 1 - i])
	}
	fmt.Println(reverse_a)
}