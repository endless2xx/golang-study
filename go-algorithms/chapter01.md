# 第二章: 算法基础

[TOC]

## 1. 插入排序

**插入排序** 对于少量元素的排序，这是一个有效的排序算法

![1561647752867](D:\Endless1905.github\go-space\src\Introduction_to_Algorithms\chapter01-insertion-sort\chapter01.assets\1561647752867.png)

**案例代码**

```go
/**
 * 插入排序方法，参数是数组的指针
 * 将牌分成三部分
 * 0 ~ (j-1) 左手中排好序的牌
 * j 当前抓到的牌，需要插入到有序数组中
 * (j+1) ~ len(ptr)  桌子上剩下的牌
 */
func insertionSort(ptr *[10]int) {
    
	for j := 1; j < len(ptr); j++ {
		key := ptr[j]
		// 将抓到的牌和手中的牌比较（一张张的比较，直到找到合适的位置）
		for i := j-1; i>=0; i-- {
			// 抓到的牌key 小于当前正在计较的牌，则将当前循环到的牌往后移一位
			if ptr[i] > key {
				ptr[i+1] = ptr[i]
				ptr[i] = key
			} else {
				break
			}
		}
	}
}
```

**循环不变式（Loop invariants）和插入排序的正确性**

**循环不变式** 子数组   **A[0, j]** 的元素其实就是 **原数组** 中的前 **[0, j]** 中的元素，只是现在是有序的了（即：你手里抓到的牌就是桌牌里面最上面的牌，之时到了你手里被你排序了）



## 2. Analyzing algorithms



## 3. Designing algorithms



## 4. Divide-and-Conquer

