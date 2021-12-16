//岳小鹏参加golang考试，他和父亲岳不群达成承诺
//如果成绩[90,100]分时，奖励一辆BMW；
//成绩为(80，99]分时，奖励一台iphone 11 plus；
//成绩为[60，80]分时，奖励一台ipad mini;
//成绩低于60分时，奖励一个破鞋。
package main

import (
	"fmt"
)

func main() {
	reward()
}
func reward() {
	var score int
	fmt.Println("请输入岳小鹏的Golang成绩: ")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Printf("岳小鹏的Golang成绩为: %v 分, 获得奖励'一辆BMW!!!'\n", score)
	} else if score > 80 && score <= 99 {
		fmt.Printf("岳小鹏的Golang成绩为: %v 分, 获得奖励'一台iphone 11 plus!!!'\n", score)
	} else if score >= 60 && score <= 80 {
		fmt.Printf("岳小鹏的Golang成绩为: %v 分, 获得奖励'一台ipad mini!!!'\n", score)
	} else {
		fmt.Printf("岳小鹏的Golang成绩为: %v 分, 获得奖励'接受爸爸的挨打吧!!!'\n", score)
	}
}
