package utils

import "fmt"

var result float64

func Calculate(num1 float64, num2 float64, operate string) float64 {

	switch operate {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("出错了!")
	}
	return result
}
