package main

import (
	"fmt"
)


var a, b string
var i, j int
var flag bool

func main() {
	fmt.Println("Строка, в которой осуществляется поиск:")
	fmt.Scan(&a)
	fmt.Println("Искомая подстрока:")
	fmt.Scan(&b)
	for i = 0; i < len(a); i++ {
		flag = true
		for j = 0; j < len(b); j++ {
			if b[j] != a[i + j] {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
			break
		}
	}
	if !flag {
		fmt.Println(-1)
	}
}