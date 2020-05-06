package main

import (
	calc "../calc"
	"fmt"
)

func main() {
	sum := calc.Add(12, 2)
	max := calc.Max(1, 2)
	min := calc.Min(1, 2)
	fmt.Printf("合计：%d 最小值：%d 最大值：%d", sum, min, max)
}
