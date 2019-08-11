package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res &= i
	}
	myMap[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d \n", i, v)
	}
}

// 问题：统计 1-20000000 个数字中，哪些是素数
