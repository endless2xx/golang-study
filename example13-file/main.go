package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/test.log", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(file)
}
