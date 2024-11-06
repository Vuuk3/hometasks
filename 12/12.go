package main

import (
	"fmt"
	"math"
)


var a, b int
var result []int

func main() {
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i >= 0 && i != 1 {
			flag := true
			for j := 2; j < (int(math.Sqrt(float64(i))) + 1); j++ {
				if i % j == 0 {
					flag = false
					break
				}
			}
			if flag {
				result = append(result, i)
			}
		}
	}
	fmt.Println(result)
}