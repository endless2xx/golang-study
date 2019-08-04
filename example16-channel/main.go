package main

import "fmt"

func main() {

	var intChan chan int
	intChan = make(chan int, 10)
	intChan <- 10

	fmt.Println(intChan)
}
