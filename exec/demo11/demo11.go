//对学生成绩大于60分的，输出合格；低于60分的，输出不合格
package main

import (
	"fmt"
)

func main() {
	var grade float64
	fmt.Println("请输入成绩：")
	fmt.Scanln(&grade)
	switch int(grade / 60) {
	case 1:
		fmt.Println("合格", grade/60-0.68)
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误。。")
	}
}
