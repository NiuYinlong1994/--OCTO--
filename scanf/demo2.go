//编写一个程序，可以输入人的年龄，如果该同志的年龄大于18，则输出你的年龄大于18，要对自己的行为负责；否则输出你的年龄小于18，请注意你的言行举止！
package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄.....")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("你的年龄大于18，要对自己的行为负责!")
	} else {
		fmt.Println("你的年龄小于18，请注意你的言行举止！")
	}
}
