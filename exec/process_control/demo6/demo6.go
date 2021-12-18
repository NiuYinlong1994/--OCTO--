//编写程序，根据输入的月份和年份，求出该月的天数（1-12月）【switch题】
//1/3/5/7/8/10/12月=========31天
//2月=======================28/29天
//其他======================30天
package main

import (
	"fmt"
)

func main() {
	var month, year int
	fmt.Println("请输入年份: ")
	fmt.Scanln(&year)
	fmt.Println("请输入月份: ")
	fmt.Scanln(&month)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		//闰年。2月有29天
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Printf("%d年%d月有31天\n", year, month)
		case 2:
			fmt.Printf("%d年%d月有29天\n", year, month)
		default:
			fmt.Printf("%d年%d月有30天\n", year, month)
		}
	} else {
		//非闰年。2月有28天
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Printf("%d年%d月有31天\n", year, month)
		case 2:
			fmt.Printf("%d年%d月有28天\n", year, month)
		default:
			fmt.Printf("%d年%d月有30天\n", year, month)
		}
	}
}
