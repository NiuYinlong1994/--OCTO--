//求一个数组的和以及平均值，for range
package main

import (
	"fmt"
)

func SumArr() {
	var Array = [...]int{32, 34, 121, 433, 454}
	var arrSum int
	for _, value := range Array {
		arrSum += value
	}
	fmt.Printf("数组中所有数之和为%d; 平均值为%0.2f.\n", arrSum, float64(arrSum)/float64(len(Array)))
}
func main() {
	SumArr()
}
