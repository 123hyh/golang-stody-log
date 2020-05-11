package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写锁(用于读的次数大于写的次数)
*/
var (
	x     int64
	wg    sync.WaitGroup
	mutex sync.RWMutex
)

func read() {
	//mutex.Lock()
	mutex.RLock()
	time.Sleep(time.Microsecond)
	//mutex.Unlock()
	mutex.RUnlock()
	wg.Done()
}
func write() {
	mutex.Lock()
	x += 1
	time.Sleep(time.Microsecond * 100)
	mutex.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
