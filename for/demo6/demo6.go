//打印金字塔经典案例
//使用for循环完成下面的案例
//请编写一个程序，可以接收整数表示层数，打印出金字塔。
//*						 //		*
//**				    //	   ***
//*** 		改进===>>	//     *****
//****				   //    *******
//*****				   //   *********

package main

import (
	"fmt"
)

func main() {
	var layerCount int
	fmt.Println("输入金字塔的层数: ")
	fmt.Scanln(&layerCount)
	fmt.Printf("下面展示%d层金字塔\n", layerCount)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------")

	for i := 1; i <= layerCount; i++ {
		//打印空格
		for k := 1; k <= layerCount-i; k++ {
			fmt.Printf(" ")
		}
		//打印*
		for j := 1; j <= (2*i - 1); j++ {
			if j == 1 || j == 2*i-1 || i == layerCount {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
