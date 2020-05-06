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

	arr7 := [2]int{1, 2}
	arr8 := arr7[0:]
	arr7[0] = 10
	fmt.Printf("%v\n", arr8)
	makeSlice()

}
func makeSlice() {
	//	mack 方法创建 切片

	arr1 := make([]int, 5 /* 长度 */, 10 /* cap */)
	arr2 := []int{1, 2, 3}
	//追加元素
	arr1 = append(arr1, arr2...)
	fmt.Printf("%v\n", arr1)
	//	copy 切片
	arr3 := make([]int, 3, 3)
	copy(arr3, arr2)
	fmt.Printf("%v\n", arr3)
	//	删除切片某个下标的值
	arr3 = append(arr3[:1], arr3[2:]...)
	fmt.Printf("%v\n", arr3)
	arr4 := [...]int{1, 2, 3, 4}
	arr5 := append(arr4[:])
	//arr5 := append(arr4[:2],arr4[3:]...)
	fmt.Printf("%v\n", arr5)
	tody()
}
func tody() {
	fmt.Printf("--------------切片练习-------------\n")

	arr1 := [...]int{1, 2, 3}

	arr2 := append(arr1[:])
	fmt.Printf("数组转切片%v\n", arr2)

	arr2 = append(arr2[:2], arr2[3:]...)
	fmt.Printf("删除切片中的某一个元素%v\n", arr2)

	arr3 := make([]int, 5, 10)
	fmt.Printf("make创建切片 %v\n", arr3)
	arr4 := make([]int, 4, 6)
	copy(arr4, arr2)
	fmt.Printf("%v\n", arr4)

	my()
	hasHan()
}
func my() {
	var arr = [...]int{1, 2, 3}
	//arr1:=make([]int,3,3)
	arr2 := make([]int, len(arr), 5)
	copy(arr2, arr[:])
	arr2 = append(arr2, arr2[2:]...)
	fmt.Printf("%v\n", arr2)
}

