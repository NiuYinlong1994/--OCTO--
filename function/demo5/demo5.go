//请编写一个函数swap(n1 int, n2 int)，可以交换n1和n2的值
package main

import (
	"fmt"
)

func Swap(n1 *int, n2 *int) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}
func main() {
	var a, b int = 3, 5
	Swap(&a, &b)
	fmt.Println(a, b)
}
