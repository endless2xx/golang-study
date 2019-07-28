package main

import (
	"fmt"
	"time"
)

// 递归
func main() {
	// recursive(0)
	result := fibonacci(3)
	fmt.Println(result)
	result = adds(100)
	fmt.Println(result)
}

// 使用递归计算 斐波那契数列
func fibonacci(n int) int {
	if n < 1 {
		panic("error")
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 使用递归计算 1+2+3+4+...+100
func adds(n int) int {
	if n == 1 {
		return 1
	}
	return adds(n-1) + n
}

func recursive(n int) {
	if n > 5 {
		return
	}
	fmt.Println("hello world, ", n)
	time.Sleep(time.Second)
	recursive(n + 1)
}
