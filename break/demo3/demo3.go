//实现登录验证，有三次机会，如果用户名为”张无忌“,密码为“888”
//提示登录成功，否则提示还有几次机会

package main

import (
	"fmt"
)

func main() {
	var user, password string
	var count int = 3
	for i := 1; i <= count; i++ {
		fmt.Println("请输入账户: ")
		fmt.Scanln(&user)
		fmt.Println("请输入密码: ")
		fmt.Scanln(&password)
		if user == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("账户/密码输入错误,还有%v次机会\n", count-i)
		}
	}
}
