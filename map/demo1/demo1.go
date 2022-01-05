package main

import (
	"fmt"
)

func main() {
	stu := make(map[string]string, 4)
	stu["张三"] = "上海"
	stu["李四"] = "北京"
	fmt.Println(stu)
}
