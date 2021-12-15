//定义一个变量保存华氏温度，华氏温度转化为摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度。
package main

import (
	"fmt"
)

func main() {

	var fahrenheit float32 = 134.2
	var centigrade float32
	centigrade = (5.0 / 9 * (fahrenheit - 100))
	fmt.Printf("华氏温度:%vF对应的摄氏温度为:%vC\n", fahrenheit, centigrade)
}
