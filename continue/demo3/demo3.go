//某人有100000元，每经过一次路口需要缴费，规则如下：
//当现金大于50000时，每次交5%；
//当现金小于等于50000时，每次交1000；
//编程计算该人可以经过多少次路口；
//使用for + break方式完成。
package main

import (
	"fmt"
)

func main() {
	var money float32 = 100000
	// var money1 int
	var count int
	for {
		if money > 50000 {
			money -= money * 0.05
			count++
		} else if money <= 50000 && money >= 1000 {
			money -= 1000
			count++
		} else {
			break
		}
	}
	fmt.Printf("可以进过路口次数: %v; 还剩: %v元钱，小于1000元不足以前往下一个路口\n", count, money)
}
