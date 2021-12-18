//实现对三个整数进行排序，输出时按照从小到大的顺序排序
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("输入三个整数: ")
	fmt.Scanf("%v %v %v\n", &a, &b, &c)
	if a > b {
		if a > c {
			//fmt.Println("a最大")
			if b > c {
				fmt.Println(c, b, a)
			} else {
				fmt.Println(b, c, a)
			}
		} else {
			//fmt.Printf("c最大")
			fmt.Println(b, a, c)
		}
	} else if b > c {
		//fmt.Println("b最大")
		if a > c {
			fmt.Println(c, a, b)
		} else {
			fmt.Println(a, c, b)
		}
	} else {
		//fmt.Println("c最大")
		fmt.Println(a, b, c)
	}
}
