//从控制台获取用户信息。[姓名，年龄，薪资，是否通过考试]
package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var sal float64
	var isPass bool
	fmt.Println("请输入您的姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入您的年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入您的薪资:")
	fmt.Scanln(&sal)
	fmt.Println("请确认是否通过:")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名: %v\n年龄: %v\n薪资: %v\n通过状态: %v\n", name, age, sal, isPass)

	//fmt.Scanf()可以按指定格式输出
	fmt.Println("请输入您的姓名, 年龄, 薪资, 以及是否通过？") //用空格隔开
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名: %v\n年龄: %v\n薪资: %v\n通过状态: %v\n", name, age, sal, isPass)
}
