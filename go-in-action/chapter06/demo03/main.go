package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 计数器，所有的 goroutine 都会修改其值
	counter int
	// 用来等待线程结束
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		// fmt.Printf("线程 %d 退出 %d %d\n", id, count, counter)
		fmt.Println("Before")
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
		fmt.Println("After")
		// fmt.Printf("线程 %d 重新执行 %d %d\n", id, count, counter)
		value++
		counter = value
	}
}
