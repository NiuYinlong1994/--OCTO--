package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
	city string
}

func main() {
	var person Student
	person.name = "小明"
	person.age = 18
	person.city = "上海"
	fmt.Println(person)
}
