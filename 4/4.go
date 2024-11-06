package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	a, err1 := bufio.NewReader(os.Stdin).ReadString('\n')
	b, err2 := bufio.NewReader(os.Stdin).ReadString('\n')
	if err1 != nil || err2 != nil{
		fmt.Println(err1, err2)
	} else {
		a = a[0:len(a) - 2]
		b = b[0:len(b) - 2]
		a = a + " " + b
		arr := strings.Split(a, " ")
		sort.Slice(arr, func(i, j int) bool {
			n1, err1 := strconv.ParseInt(arr[i], 10, 64)
			n2, err2 := strconv.ParseInt(arr[j], 10, 64)
			if err1 == nil && err2 == nil {
				return n1 < n2
			}
			return false
		})
		fmt.Println(arr)
	}
}