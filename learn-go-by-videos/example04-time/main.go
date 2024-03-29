package main

import (
	"fmt"
	"time"
)

// go 语言中的时间类型
func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Printf("cost: %d us\n", (end-start)/1000/1000)

}

func test() {
	time.Sleep(time.Millisecond * 100)
}
