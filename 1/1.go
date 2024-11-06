package main

import (
	"fmt"
	"strconv"
)

var start, finish int
var n string
var err error
var i int64

func main() {
	fmt.Scan(&n, &start, &finish)
	i, err = strconv.ParseInt(n, start, 64)
	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println(strconv.FormatInt(i, finish))
	}
}
/*
Но я живу, не видя дня,
*/
// Во мраке бесконечной ночи