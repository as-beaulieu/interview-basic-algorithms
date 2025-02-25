package main

import "fmt"
import "strings"
import "strconv"

func ReversePolishNotation(str string) string {
	split := strings.Split(str, " ")
	var calc []int
	var result int

	for _, item := range split {

		number, err := strconv.Atoi(item)
		if err != nil {
			calc, result = performOperation(calc, item)
		} else {
			calc = append(calc, number)
		}
	}

	// code goes here
	str = strconv.Itoa(result)
	return str
}

func performOperation(calc []int, operator string) ([]int, int) {
	result := 0
	if len(calc) < 2 {
		fmt.Println("error")
		return nil, 0
	}
	i := len(calc)
	a := calc[i-2]
	b := calc[i-1]
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	var resultCalc []int
	resultCalc = append(resultCalc, calc[:len(calc)-2]...)
	resultCalc = append(resultCalc, result)

	return resultCalc, result
}

func main() {
	input := "3 4 + 7 2 + *"
	fmt.Println(ReversePolishNotation(input))

}
