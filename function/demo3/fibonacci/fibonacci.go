//斐波那契数
//使用递归的方式，求出斐波那契数1,1,2,3,5,8,13.... f(1)=1;f(2)=1;f(n)=f(n-1)+f(n-2)(n>=3)
//给你一个整数，求出它的值是多少？
package fibo

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
