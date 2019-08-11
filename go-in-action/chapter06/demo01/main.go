package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 演示如何创建 goroutine
// 以及调度器的行为
func main() {
	// 分配一个逻辑处理器给调度器
	// 函数 GOMAXPROCS 允许程序更改调度器可以使用的逻辑处理器的数量
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	// 计数加2，表示等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()
		// 打印字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Printf("\n")
		}
	}()

	fmt.Printf("\n")
	// 声明一个匿名函数，并创建第二个 goroutine
	go func() {
		// 同上
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Printf("\n")
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
