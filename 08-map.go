package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
map 类型 （key value）散列表
*/
func main() {
	m1 := make(map[string]string, 10)
	m1["age"] = "18"
	m1["name"] = "黄裕辉"

	fmt.Printf("%v", m1["age"])
	//	map 中是否存在某个key
	val, ok := m1["name"] // 如果不存在key 会拿到类型的初始值
	println(val, ok)
	//	遍历map
	for key, val := range m1 {
		println(key, val)
	}
	//	删除map的某个键
	delete(m1, "name")
	fmt.Printf("%v\n", m1)
	//	元素为map的切片
	arr2 := make([]map[string]string, 10, 100)
	arr2[0] = make(map[string]string, 10)
	arr2[0]["name"] = "huang"
	fmt.Printf("%v\n", arr2)
	//	值为切片的map
	arr3 := make(map[string][]int, 10)
	arr3["name"] = []int{1, 2, 3, 4}
	fmt.Printf("%v\n", arr3)

	arr4 := make([]map[string]string, 10, 10)
	for i := range arr4 {
		item := make(map[string]string)
		item["name"] = "黄"
		arr4[i] = item
	}
	fmt.Printf("%v\n", arr4)
	arr5 := make(map[string][]string, 10)
	arr5["城市"] = []string{"gz", "sh", "sz"}
	fmt.Printf("%v\n", arr5)
	wordNumber()
	hasHan()
	judge()
}

//单词出现的次数
func wordNumber() {
	str := "hello word how do you do"
	str1 := make(map[string]int)
	for _, val := range strings.Split(str, " ") {
		if _, ok := str1[val]; ok {
			str1[val]++
		} else {
			str1[val] = 1
		}
	}
	fmt.Printf("%v\n", str1)
}

//汉字的次数
func hasHan() {
	str := "hello黄裕辉"
	for _, val := range str {
		if unicode.Is(unicode.Han, val) {
			fmt.Printf("%v\n", string(val))
		}
	}
}

//回文判断
func judge() {
	
}
