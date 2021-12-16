//使用switch把小写类型的char型转为大写（键盘输入）
//只转换a,b,c,d,e，其他的输出"other"
package main

import (
	"fmt"
)

func main() {
	var letter byte
	fmt.Println("请输入一个char: ")
	fmt.Scanf("%c", &letter)
	switch letter {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
}
