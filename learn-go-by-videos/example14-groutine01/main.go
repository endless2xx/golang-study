package main

import (
	"fmt"
	"time"
)

func test() {
	var i int
	for {
		i++
		fmt.Printf("i in test: %v \n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go test()

	var i int
	for {
		i++
		fmt.Printf("i in main: %v \n", i)
		time.Sleep(time.Second)
	}
}
