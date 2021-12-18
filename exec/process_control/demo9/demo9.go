//有两个数a和b，如果a能够被b整除或者a+b大于1000,则输出a，否则输出b
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("请输入两个整数: ")
	fmt.Scanf("%v %v\n", &a, &b)
	if a%b == 0 || a+b > 1000 {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
