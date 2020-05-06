package main

import "fmt"

func main() {
	//数组 是 值类型
	//声明数组 【长度】值类型 {值，值}
	var arr1 [3]bool
	arr2 := [3]bool{}

	// 初始化值
	arr1 = [3]bool{true, false, true}
	// 根据初始值 推断数组的长度
	arr3 := [...]bool{}

	//根据索引来初始化值
	arr4 := [3]bool{0: true, 2: true}
	fmt.Printf("%v %v %v %v\n", arr1, arr2, arr3, arr4)

	//多维数组 [[1,2],[3,4],[5, 6]]
	arr5 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Printf("%v \n", arr5)
	mapArr()
}

/**
数组遍历
*/
func mapArr() {
	arr1 := [3]bool{0: true, 2: true}
	for _, val := range arr1 {
		fmt.Printf("for range:%v\n", val)
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("for i :%v\n", arr1[i])
	}
	//	求和
	arr2 := [...]int{1, 3, 5, 7, 8}
	var sum int = 0
	for _, val := range arr2 {
		sum += val
	}
	fmt.Printf("求和 %v \n", sum)

	//	找出和 为 8 的 元素下标
	for i, val := range arr2 {
		for i2, val2 := range arr2 {
			if val+val2 == 8 && i2 > i {
				fmt.Printf("和为8的元素下标 %v %v \n",i, i2)
			}
		}
	}
}
