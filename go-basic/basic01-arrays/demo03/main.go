package main

// 使用指针传递数组
func main() {
	array := [100]int{}
	foo(&array)
}

func foo(array *[100]int) {
	// ...
}
