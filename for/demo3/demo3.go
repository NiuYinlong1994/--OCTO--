//完成下面表达式的输出,6是可变的
// 0 + 6 = 6
// 1 + 5 = 6
// 2 + 4 = 6
// 3 + 3 = 6
// 4 + 2 = 6
// 5 + 1 = 6
// 6 + 0 = 6
package main

import (
	"fmt"
)

func main() {
	var sum int = 6
	for i := 0; i <= sum; i++ {
		fmt.Printf("%d + %d = %d\n", i, sum-i, i)
	}
}
