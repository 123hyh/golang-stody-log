package main

import (
	calc "../calc"
	"fmt"
)

func main() {
	sum := calc.Add(12, 2)
	fmt.Printf("%d\n", sum)
}
