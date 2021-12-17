//编写一个程序，可以打印10句
//“你好～尚硅谷～”
package main

import (
	"fmt"
)

func main() {
	var helloString string = "你好～尚硅谷～"
	for i := 0; i < 10; i++ {
		fmt.Println(helloString)
	}
}
