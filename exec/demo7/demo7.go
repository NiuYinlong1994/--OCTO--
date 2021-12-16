//求ax2+bx+c=0方程的根。a,b,c分别为函数的参数，如果b2-4ac>0，则有两个解；b2-4ac=0，则有一个解；b2-4ab<0，则无解。
//提示1：x1=(-b+sqrt(b2-4ac))/2a;
//      x2=(-b-sqrt(b2-4ac))/2a;
//提示2: math.Sqrt(num); 需要引入math包；
//测试数据：3,100,6
package main

import (
	"fmt"
	"math"
)

func equationMath() {
	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0
	x := b*b - 4*a*c
	if x > 0 {
		fmt.Printf("次方程有两个解x1=%v,x2=%v\n", (-b+math.Sqrt(x))/2*a, (-b-math.Sqrt(x))/2*a)
	} else if x == 0 {
		fmt.Printf("次方程有两个相同的解x1,x2=%v\n", (-b+math.Sqrt(x))/2*a)
	} else {
		fmt.Println("次方程无解")
	}
}
func main() {
	equationMath()
}
