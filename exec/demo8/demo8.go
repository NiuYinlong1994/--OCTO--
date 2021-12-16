//参加百米运动会，如果用时8秒以内的进入决赛
//否则提示淘汰，并且根据性别提示进入男子组或女子组
//输入成绩和性别
package main

import (
	"fmt"
)

func sport() {
	var time float64
	fmt.Println("请输入成绩: ")
	fmt.Scanln(&time)
	if time <= 8 {
		var sex string
		fmt.Println("请输入性别: ")
		fmt.Scanln(&sex)
		if sex == "男" {
			fmt.Println("恭喜! 进入男子组")
		} else {
			fmt.Println("恭喜! 进入女子组")
		}
	} else {
		fmt.Println("淘汰...")
	}
}
func main() {
	sport()
}
