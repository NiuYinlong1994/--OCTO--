//函数可以没有返回值案例，编写一个函数，从终端输入一个整数答应出对应的金字塔
//    *
//   ***
//  *****
// *******
//*********
package main

import (
	"fmt"
)

func Pyramid(layerCount int) {
	for i := 1; i <= layerCount; i++ {
		for k := 1; k <= layerCount-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	Pyramid(6)
}
