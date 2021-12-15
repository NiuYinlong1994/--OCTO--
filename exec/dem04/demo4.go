//求两个数的最大值
//求三个数的最大值
package main

import (
	"fmt"
)

//求两个数a和b的最大值
func TwoMax() {
	var a int = 67
	var b int = 89
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

//求三个数c,d,e的最大值
func ThreeMax() {
	var c int = 18971000
	var d int = 1786200
	var e int = 181760
	if c > d {
		if c > e {
			fmt.Println(c)
		} else {
			fmt.Println(e)
		}
	} else if d > e {
		fmt.Println(d)
	} else {
		fmt.Println(e)
	}
}
func main() {
	ThreeMax()
	TwoMax()
}
