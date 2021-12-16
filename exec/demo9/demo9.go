//出票系统：根据淡旺季的月份和年龄，打印票价（优先考虑学生先做）
//4～10月份是旺季：
//		成人（18～60岁）：60元/人
//		儿童（小于18岁）：半价（30元/人）
//		老人（大于60岁）：1/3（20元/人）
//淡季：
//		成人：40元/人
//		其他：20元/人

package main

import (
	"fmt"
)

func ticketPrice() {
	var season int
	var age int
	fmt.Println("请输入月份: ")
	fmt.Scanln(&season)
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)
	if season >= 4 && season <= 10 { //旺季
		if age >= 18 && age < 60 {
			fmt.Println("旺季成人的票价是60元/人")
		} else if age > 0 && age < 18 {
			fmt.Println("旺季儿童的票价是30元/人")
		} else if age >= 60 {
			fmt.Println("旺季老人的票价是20元/人")
		}
	} else { //淡季
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人的票价是40元/人") //成人（18～60岁）
		} else {
			fmt.Println("淡季其他人的票价是20元/人") //其他
		}
	}
}
func main() {
	ticketPrice()
}
