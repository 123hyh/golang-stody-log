package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 50000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main() {
	wg.Add(2)

	go add()
	go add()

	fmt.Println(x)
	wg.Wait()
}
