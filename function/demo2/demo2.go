//求函数值
//已知f(1)=3;f(n)=2*f(n-1)+1
//请使用递归的思想编程，求出f(1)的值
package main

import (
	"fmt"
)

func Sum(x int) int {
	if x == 1 {
		return 3
	} else {
		return Sum(x-1)*2 + 1
	}
}
func main() {
	var num int = 2
	result := Sum(num)
	fmt.Println("结果", result)
}
