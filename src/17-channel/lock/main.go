package main

/**
互斥锁(多个goroutine 同时设置一个变量时)
*/
import (
	"fmt"
	"sync"
)

var (
	swg  sync.WaitGroup
	y    int64 = 0
	lock sync.Mutex
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 枷锁
		y += 1
		lock.Unlock() // 解锁
	}
	swg.Done()
}

func main() {
	swg.Add(2)

	go add()
	go add()
	swg.Wait()

	fmt.Println(y)
}
