package main

import "fmt"

// 闭包
// 把返回的函数想象成一个对象
// 变量 x 就是该对象中的一个属性
func main() {

	f := AddUpper()
	fmt.Println(f(10))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

// 闭包：累加器
func AddUpper() func(int) int {
	var n int = 0
	return func(x int) int {
		n += x
		return n
	}
}
