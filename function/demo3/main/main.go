package main

import (
	"fmt"
	fibo "utils/fibonacci"
)

func main() {
	var num int = 3
	result := fibo.Fibonacci(num)
	fmt.Println(result)
}
