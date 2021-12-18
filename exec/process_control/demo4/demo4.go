//写出输出结果
/*
m,n :=0, 3
if (m>0){
	if (n>2){
		fmt.Println("A")
	}else {
		fmt.Println("B")
	}
}
*/
package main

import (
	"fmt"
)

func main() {
	m, n := 0, 3
	if m > 0 {
		if n > 2 {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	}
}

//啥也不输出
