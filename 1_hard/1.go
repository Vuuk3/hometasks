package main

import (
	"fmt"
	"strconv"
)


var s, u0, a, t string


func main() {
	fmt.Scan(&u0, &a, &t)
	u0, err1 := strconv.ParseFloat(u0, 64)
	a, err2 := strconv.ParseFloat(a, 64)
	t, err3 := strconv.ParseFloat(t, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Ошибка ввода данных")
	} else {
		fmt.Println(u0 * t + 0.5 * a * t * t)
	}
}