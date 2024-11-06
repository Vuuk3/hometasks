package main

import (
	"fmt"
	"math"
	"strconv"
)


var a, b, check_sum uint64
var result []uint64

func main() {
	fmt.Scan(&a, &b)
	for i := uint64(a); i <=uint64(b); i++ {
		r := strconv.FormatUint(i, 10)
		check_sum = 0
		for j := 0; j < len(r); j ++ {
			n, err := strconv.ParseUint(string(r[j]), 10, 64)
			if err == nil {
				check_sum += uint64(math.Pow(float64(n), float64(len(r))))
			}
		}
		if check_sum == uint64(i) {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}