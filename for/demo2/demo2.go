//打印1～100之间所有是9的倍数的整数的个数及总和
package main

import (
	"fmt"
)

func main() {
	var sum, count int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			sum += i
			count++
		}
	}
	fmt.Printf("9的倍数的整数的个数: %d, 总和: %d\n", count, sum)
}
