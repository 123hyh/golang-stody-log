package main

import "fmt"

/**
切片类型 引用类型
*/
func main() {
	//切片的定义
	var arr []int
	arr1 := []int{1, 2}
	fmt.Printf("(%v %T) \n", arr1, arr)

	//	nul 是否开辟了内存空间
	fmt.Printf("%v\n", arr == nil)

	//	切片的长度和容量
	fmt.Printf("数组的长度和容量(%d, %d)\n", len(arr1), cap(arr1))

	//	由数组得到切片
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7}

	arr3 := arr2[0:]
	arr4 := arr2[2:4]
	fmt.Printf("基于数组得到切片%v %v \n", arr3, arr4)

	arr5 := []int{1, 2, 3, 4}
	fmt.Printf("%v\n", arr5)
	arr6 := arr5[0:]
	arr5[1] = 20
	fmt.Printf("%v\n", arr6)

}
