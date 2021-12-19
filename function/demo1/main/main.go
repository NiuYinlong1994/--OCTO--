package main

import (
	"fmt"

	"utils"
)

func main() {
	var num1, num2 float64
	var operate string
	fmt.Println("请输入两个数: ")
	fmt.Scanf("%v %v\n", &num1, &num2)
	fmt.Println("请输入运算符: ")
	fmt.Scanln(&operate)
	res := utils.Calculate(num1, num2, operate)
	fmt.Printf("%v %v %v = %.2f\n", num1, operate, num2, res)
}
