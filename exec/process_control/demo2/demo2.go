//判断一个年份是否为闰年？
//能够被4整除并且不能被100整除，或者能被400整除。
package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Println("请输入年份: ")
	fmt.Scanln(&year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%d是闰年\n", year)
	} else {
		fmt.Printf("%d不是闰年\n", year)
	}
}
