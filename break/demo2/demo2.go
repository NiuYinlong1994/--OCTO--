//100以内的数求和，求出和为第一次大于20的当前数
package main

import (
	"fmt"
)

func main() {
	var numSum int = 0
	for i := 1; i <= 100; i++ {
		numSum += i
		if numSum > 20 {
			fmt.Println(i, numSum)
			break
		}
	}
}
