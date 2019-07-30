package main

import "fmt"

// 闭包
// 把返回的函数想象成一个对象
// 变量 x 就是该对象中的一个属性
func main() {

	f := test()
	fmt.Println(f(10))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

func test() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
