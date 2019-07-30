package main

import (
	"fmt"
	"strings"
)

/*
	没有闭包的时候，函数就是一次性买卖，函数执行完毕后就无法再更改函数中变量的值（应该是内存释放了）；
	有了闭包后函数就成为了一个变量的值，只要变量没被释放，函数就会一直处于存活并独享的状态，
	因此可以后期更改函数中变量的值（因为这样就不会被go给回收内存了，会一直缓存在那里）。
*/
func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("dog"))
	fmt.Println(f("tom.jpg"))

	f = makeSuffix(".avi")
	fmt.Println(f("cats"))
	fmt.Println(f("fish.avi"))
}

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		// 如果没有后缀，则添加
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
