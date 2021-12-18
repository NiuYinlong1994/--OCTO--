//开发一款软件，输入您的身高和体重。根据公式(身高-108)*2=体重，可以有10斤左右的浮动。来观察测试者是否符合？
package main

import (
	"fmt"
)

func main() {
	var height, weight, weight1 float64
	fmt.Println("请输入您的身高: ")
	fmt.Scanln(&height)
	fmt.Println("请输入您的体重: ")
	fmt.Scanln(&weight)
	weight1 = (height - 108) * 2
	if weight <= weight1-10 {
		fmt.Printf("您的身高为%vcm; 您的体重是%v; 体脂偏低\n", height, weight)
	} else if weight >= weight1+10 {
		fmt.Printf("您的身高为%vcm; 您的体重是%v; 体脂偏高\n", height, weight)
	} else {
		fmt.Printf("您的身高为%vcm; 您的体重是%v; 体脂健康\n", height, weight)
	}
}
