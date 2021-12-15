package main

import (
	"fmt"
)

func main() {
	var a int = 300
	var b int = 400
	var ptr *int = &a       //ptr指向a变量的地址
	*ptr = 200              //赋值ptr,相当于修改变量a的值
	ptr = &b                //重新把ptr指向b变量的地址
	*ptr = 500              //赋值ptr,相当于修改变量b的值
	fmt.Println(a, b, *ptr) //200,500,500
}
