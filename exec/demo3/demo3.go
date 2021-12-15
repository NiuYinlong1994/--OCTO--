//有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("交换前两变量的值分别为：%d,%d\n", a, b)
	a = a + b
	b = a - b // b = a + b - b = a
	a = a - b // a = a + b - a = b
	fmt.Printf("交换后两变量的值分别为：%d,%d\n", a, b)
}
