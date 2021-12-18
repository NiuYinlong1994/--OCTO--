//判断一个整数是否是水仙花数？？
//所谓水仙花数是指一个3位数，其各个位上数字立方和等于其本身，
//例如：153 = 1*1*1 + 5*5*5 + 3*3*3
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("请输入一个整数: ")
	fmt.Scanln(&num)
	//判断该数为一个三位数，如果
	if num > 999 || num < 100 {
		fmt.Println("输入错误...")
	} else if (num/100)*(num/100)*(num/100)+(num/10)%10*(num/10%10)*(num/10%10)+(num%10)*(num%10)*(num%10) == num {
		fmt.Printf("%d是一个水仙数\n", num)
	} else {
		fmt.Printf("%d不是一个水仙数\n", num)
	}
}
