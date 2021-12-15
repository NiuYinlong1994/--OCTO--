//1)写一个程序，获取一个int变量num的地址，并显示到终端

package main

import (
	"fmt"
)

// func main(){
// 	var num int
// 	fmt.Println(&num)

// }

//2)将num的地址赋给指针ptr，并通过ptr去修改num的值

func main() {
	var num int = 9
	//1)写一个程序，获取一个int变量num的地址，并显示到终端
	fmt.Println(&num)

	//2)将num的地址赋给指针ptr，并通过ptr去修改num的值
	var ptr *int = &num
	*ptr = 100
	fmt.Println(*ptr, num)
}
