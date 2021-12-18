//判断次考试成绩为什么等级？【switch题和if else if多分支题】
//90～100分之间优秀
//80～89分之间优良
//70～79分之间良好
//60～69分之间及格
//60分以下不及格

package main

import (
	"fmt"
)

func main() {
	var score float64
	fmt.Println("输入成绩: ")
	fmt.Scanln(&score)
	if score >= 90 && score <= 100 {
		fmt.Println("优秀")
	} else if score >= 80 && score < 90 {
		fmt.Println("优良")
	} else if score >= 70 && score < 80 {
		fmt.Println("良好")
	} else if score >= 60 && score < 70 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
