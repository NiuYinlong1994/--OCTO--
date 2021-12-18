//保存用户名和密码，判断用户名是否为“张三”，密码是否为“1234”，
//如果是，提示登录成功，否则提示登录失败
package main

import (
	"fmt"
)

func main() {
	var user, password string
	fmt.Println("请输入您的账户: ")
	fmt.Scanln(&user)
	fmt.Println("请输入您的密码: ")
	fmt.Scanln(&password)
	if user == "张三" && password == "1234" {
		fmt.Println("登陆成功!")
	} else {
		fmt.Println("登陆失败!")
	}

}
