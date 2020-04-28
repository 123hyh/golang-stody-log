package main

import "fmt"

func main() {
	var n1 uint8 = 1
	const n2 uint8 = 2
	fmt.Println(n1, n2)

	keywordIota()
}

/**
*  iota关键字 (仅限于在 const 中使用 ，！！！每换行一次自动 + 1)
 */
func keywordIota() {
	const (
		s1 = iota
		_ = 1
		s2 = iota
		_ = 3
		s3 =iota
		s4
	)
	println(s1, s2, s3, s4)
}
