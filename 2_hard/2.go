package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода данных")
	} else {
		input = input[:len(input) - 2]
		array := strings.Split(input, " ")
		indexOfNumber1 := 0
		indexOfNumber2 := 0
		for i := 1; i < len(array); i++ {
			i -= 3
			switch (array[i + 3]) {
			case "*":
				number1, err1 := strconv.ParseFloat(array[indexOfNumber1], 64)
				number2, err2 := strconv.ParseFloat(array[indexOfNumber2], 64)
				if err1 == nil && err2 == nil {
					array = append(array[:indexOfNumber2], array[indexOfNumber2 + 1:]...)
					array = append(array[:indexOfNumber1], array[indexOfNumber1 + 1:]...)
					array = append(append(array[:i - 2], strconv.FormatFloat(number1 * number2, 'g', -1, 64)), array[i - 1:]...)
				}
			case "+":
				number1, err1 := strconv.ParseFloat(array[indexOfNumber1], 64)
				number2, err2 := strconv.ParseFloat(array[indexOfNumber2], 64)
				if err1 == nil && err2 == nil {
					array = append(array[:indexOfNumber2], array[indexOfNumber2 + 1:]...)
					array = append(array[:indexOfNumber1], array[indexOfNumber1 + 1:]...)
					array = append(append(array[:i - 2], strconv.FormatFloat(number1 + number2, 'g', -1, 64)), array[i - 1:]...)
				}
			case "-":
				number1, err1 := strconv.ParseFloat(array[indexOfNumber1], 64)
				number2, err2 := strconv.ParseFloat(array[indexOfNumber2], 64)
				if err1 == nil && err2 == nil {
					array = append(array[:indexOfNumber2], array[indexOfNumber2 + 1:]...)
					array = append(array[:indexOfNumber1], array[indexOfNumber1 + 1:]...)
					array = append(append(array[:i - 2], strconv.FormatFloat(number1 - number2, 'g', -1, 64)), array[i - 1:]...)
				}
			case "/":
				number1, err1 := strconv.ParseFloat(array[indexOfNumber1], 64)
				number2, err2 := strconv.ParseFloat(array[indexOfNumber2], 64)
				if err1 == nil && err2 == nil {
					array = append(array[:indexOfNumber2], array[indexOfNumber2 + 1:]...)
					array = append(array[:indexOfNumber1], array[indexOfNumber1 + 1:]...)
					array = append(append(array[:i - 2], strconv.FormatFloat(number1 / number2, 'g', -1, 64)), array[i - 1:]...)
				}
			default:
				i += 3
				indexOfNumber1, indexOfNumber2 = indexOfNumber2, i
			}
		}
		fmt.Println(array[0])
	}
}