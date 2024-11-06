package main

import (
	"fmt"
)


var n int

func main() {
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Таких чисел нет")
	} else if n == 0 {
		fmt.Println(0)
	} else if n == 1 {
		fmt.Println(0)
		fmt.Println(1)
	} else {
		n1 := 0
		n2 := 1
		fmt.Println(0)
		for i := 0; (n1 + n2) < n; i++ {
			n1, n2 = n1 + n2, n1
			fmt.Println(n1)
		}
	}
}