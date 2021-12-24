//求出一个数组的最大值，并打印出对应的下标
package main

import (
	"fmt"
)

func MaxValue() {
	var num = [...]int{99, 89, 1, 2, 3, 4, 5, 6, 22, 2, 1, 2123, 4323, 54654, 765756, 85786, -1, 222222222, 232, 121, 33}
	maxvalue := num[0]
	maxindex := 0
	for i := 1; i < len(num); i++ {
		if maxvalue < num[i] {
			maxvalue = num[i]
			maxindex = i
		}
	}
	fmt.Println(maxindex, maxvalue)
}
func main() {
	MaxValue()
}
