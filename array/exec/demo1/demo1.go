//创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。使用for循环访问所有元素并打印出来
//提示：字符数据运算'A'+1='B'

package main

import (
	"fmt"
)

func Letter() {
	var letterArr [26]byte
	for i := 0; i < 26; i++ {
		letterArr[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c \n", letterArr[i])
	}
}
func main() {
	Letter()
}
