package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var answer = ""

func main() {
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		a = a[0:len(a) - 2]
		a = strings.ReplaceAll(a, ".", "")
		a = strings.ReplaceAll(a, ",", "")
		a = strings.ReplaceAll(a, "!", "")
		a = strings.ReplaceAll(a, "?", "")
		a = strings.ReplaceAll(a, "'", "")
		arr := strings.Split(a, " ")
		for i := 0; i < len(arr); i++ {
			if len(arr[i]) > len(answer) {
				answer = arr[i]
			}
		}
		fmt.Println(answer)
	}
}