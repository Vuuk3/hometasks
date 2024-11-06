package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		a = a[0:len(a) - 2]
		arr := strings.Split(a, " ")
		sort.Slice(arr, func(i, j int) bool {
			n1, err1 := strconv.ParseInt(arr[i], 10, 64)
			n2, err2 := strconv.ParseInt(arr[j], 10, 64)
			if err1 == nil && err2 == nil {
				return math.Abs(float64(n1)) < math.Abs(float64(n2))
			}
			return false
		})
		fmt.Println(arr)
	}
}