package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 字符串转 rune 类型
	var s1 string = "z罗湖/*_q`"

	var s2 = []rune(s1)
	s3 := []rune("")
	var s4 = byte('1')
	fmt.Printf("rune类型：%T\n", s3, s4)

	for _, val := range s2 {
		if val > 122 {
			s3 = append(s3, val)
		}
	}
	//	int 转 string
	s5 := strconv.Itoa(1)

	fmt.Printf("int 转 string类型：%v\n", s5)


	fmt.Printf("s1的汉字总长度：%v\n", len(s3))
	fmt.Printf("rune类型：%T\n", s3)

	// int 转 flot
	var n1 int = 1
	var n2 float32 = 3.12
	fmt.Printf("int转float：%v", float32(n1)+n2)

}
