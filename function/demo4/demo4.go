//有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个，以后每天猴子都吃其中的一半，然后再多吃一个；
//当到第10天在想吃时，发现只有一个桃子了，问题：最初共有多少个桃子？
//思路分析：
//第10天只有1个桃子
//第9天有几个桃子？  ==>  (第10天的桃子数+1)*2
//规律：第n天的桃子数据：peach(n) = (peach(n+1)+1)*2

package main

import (
	"fmt"
)

func Peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (Peach(n+1) + 1) * 2
	}
}
func main() {
	num := Peach(1)
	fmt.Printf("第10天有%v个桃子\n", num)
}
