//请编写一个程序，具体如下：
//编写一个函数makeSuffix(suffix string)可以接收一个文件后缀名(比如: .jpg)，并返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如: .jpg), 则返回文件名.jpg，如果已经有.jpg后缀，则返回原文件名
//要求使用闭包的方式完成
//strings.HasSuffix

package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(myString string) string {
		//如果myString没有指定的后缀，则加上；否则返回原来的名字
		if !strings.HasSuffix(myString, suffix) {
			return myString + suffix
		}
		return myString
		// if strings.HasSuffix(myString, suffix) == false {
		// 	return myString + suffix
		// } else {
		// 	return myString
		// }
	}
}
func main() {
	myName := makeSuffix(".jpg")
	fmt.Println(myName("go.jpg"))
	fmt.Println(myName("python"))
	fmt.Println(myName("java.jpg"))
	fmt.Println(myName("php"))
}
