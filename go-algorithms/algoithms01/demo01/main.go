package main

import "fmt"

// 排序算法之插入排序

func main() {
	arr := [6]int{3, 5, 4, 1, 2, 6}
	insertionSort(&arr)
	fmt.Println(arr)
}

// 插入排序算法
// @param arr 指向待排序的数组的指针
func insertionSort(arr *[6]int) {

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := 0
		for i = j - 1; i >= 0; i-- {
			if arr[i] > key {
				arr[i+1] = arr[i]
				arr[i] = key
			} else {
				break
			}
		}
	}
}
