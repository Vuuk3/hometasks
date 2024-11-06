package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		a = a[0:len(a) - 2]
		a = strings.ReplaceAll(a, " ", "")
		a = strings.ReplaceAll(a, "?", "")
		a = strings.ReplaceAll(a, ".", "")
		a = strings.ReplaceAll(a, ",", "")
		a = strings.ReplaceAll(a, "'", "")
		a = strings.ReplaceAll(a, "!", "")
		a = strings.ToLower(a)
		flag := true
		for i := 0; int(i) < (len(a) - 1) / 2 + 1; i++ {
			if a[i] != a[len(a) - 1 - int(i)] {
				flag = false
				break
			}
		}
		fmt.Println(flag)
	}
}