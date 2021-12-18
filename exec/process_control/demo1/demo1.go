//实现判断一个整数属于那个范围？大于0？小于0？等于0
package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("请输入任意一个整数: ")
	fmt.Scanln(&number)
	if number > 0 {
		fmt.Println("该整数大于0")
	} else if number == 0 {
		fmt.Println("该整数等于0")
	} else {
		fmt.Println("该整数小于0")
	}
}
