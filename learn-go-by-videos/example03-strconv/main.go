package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将整型数组转成字符串
	result := strconv.Itoa(1000)
	fmt.Println("整型数字转字符串：", result)

	// 将字符串类型的数字，转成整型数字
	str := "2000"
	intResult, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("can't convert to int, ", err)
	}
	fmt.Println("字符串装整型数字: ", intResult+1)
}
