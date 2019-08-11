package main

import (
	"errors"
	"fmt"
)

/**
 * Go 语言中的内置函数，不需要引入包就能使用的函数
 * 1. close 用来关闭 channel
 * 2. len 用来计算长度，比如：string, array, slice, map, channel
 * 3. new 用来分配内存，主要用来分配 `值类型`，比如 int, struct 返回的是指针
 * 4. make 用来分配内存，主要用来分配 `引用类型`，比如 chan, map, slice
 * 5. append 用来追加元素到 数组中(array)或切片中(slice)
 * 6. panic 和 recover 用来做错误处理
 */

func main() {
	// **********************************************************
	// 内置函数 new 的使用，使用 new 方法是返回的是内存中的地址(指针)
	var i int
	fmt.Println(i)
	j := new(int) // 返回内存中的地址(指针)
	fmt.Println(j)
	*j = 1000       // 为该地址的内存赋值
	fmt.Println(*j) // 输出该内存地址上的值

	// new 和 make 的区别
	// new 出来的是一个空的 slice，需要初始化后才能使用
	s1 := new([]int)
	fmt.Println(s1)  // output: &[]
	fmt.Println(*s1) // output: []
	*s1 = make([]int, 5)
	(*s1)[0] = 100
	(*s1)[1] = 101
	fmt.Println(*s1) // output: []

	s2 := make([]int, 3)
	fmt.Println(s2) // output: [0 0 0]

	// **********************************************************

	// for i := 0; i < 5; i++ {
	// 	test(i)
	// }

	// **********************************************************
	// 使用 panic 处理错误信息，相当于抛出异常
	// err := initConfig()
	// if err != nil {
	// 	panic(err)
	// }
}

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test(i int) {
	// 此处捕获了系统抛出的异常，但程序还会继续往下执行
	// 否则程序将会终止，后面的程序将无法执行
	// **********************************************************
	// 使用 recover 处理异常，相当于捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	a := 100 / i
	fmt.Println(a)
	return
}
