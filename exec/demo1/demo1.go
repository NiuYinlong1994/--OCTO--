//假如还有97天放假，问：还剩多少个星期多少天？？
package main

import (
	"fmt"
)

func main() {
	var sumDay int = 97
	var week int
	var day int
	week = sumDay / 7
	day = sumDay % 7
	fmt.Printf("还剩%d个星期%d天\n", week, day)
}
