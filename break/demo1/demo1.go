//随机生成1-100的数，直到生成99这个数，看一看你一共用了多少次？
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int

	for {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100) + 1
		fmt.Println(randNum)
		count++
		if randNum == 99 {
			break
		}
	}
	fmt.Println(count)
}
