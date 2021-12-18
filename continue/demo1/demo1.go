//continue实现，打印1～100之内的奇数（要求使用for循环+continue）
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v是一个奇数\n", i)
	}
}
