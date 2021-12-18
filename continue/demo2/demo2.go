//从键盘读入个数不确定的整数，并判读读入的正数和负数的个数，输入为0时结束程序
//使用for循环+break+continue完成。positive正数 negative负数

package main

import (
	"fmt"
)

func main() {
	var num int
	var positiveCount, negativeCount int
	for {
		fmt.Println("输入一个整数: ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数的个数: %v; 负数的个数: %v \n", positiveCount, negativeCount)
}
