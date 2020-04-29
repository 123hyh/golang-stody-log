package main

import (
	"fmt"
	"strconv"
)

/**
判断
*/
func main() {
	name := "hyh"
	//带作用域 的 age ，仅限在 if 内使用
	if age := 19; age >= 18 {
		fmt.Printf("%v\n", `大于18岁`)
	} else if age < 18 {
		fmt.Printf("%v\n", `小于于18岁`)
	} else if name == "hyh" {
		fmt.Printf("%v\n", `青少年`)
	}
	fors()
}

/**
for 循环预防
*/
func fors() {

	//	1
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", i)
	}

	// 	2
	i := 0
	for i < 2 {
		fmt.Printf("%v\n", i)
		i++
	}

	forRange()

}

// 	3	for range

func forRange() {
	var s1 = "hello word 南山"
	s2 := ""
	for _, val := range s1 {
		s2 += string(val)
		//fmt.Printf("%d %c\n", i, val)
	}
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", multiplicationTable(1))

}

//打印 乘法表
var table = ``

func multiplicationTable(n int) string {
	if n <= 9 && n > 0 {
		for i := 1; i <= n; i++ {
			table += strconv.Itoa(n) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(i*n) + "; "
		}
		table += "\n"
		multiplicationTable(n + 1)
	}
	return table
}
