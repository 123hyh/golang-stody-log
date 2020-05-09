package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// 通道
var ch = make(chan int,1) // 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道

func hello(i int) {
	ch <- i
	fmt.Println(`hello~`, `发送成功`, i)
	wg.Done()
}

// 并发 goroutine
func main() {

	runtime.GOMAXPROCS(12) // 设置多核执行任务

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go hello(i)
	}

	//匿名函数 并非
	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func(i int) {
			//接收通道的值
			chalVal := <-ch
			fmt.Println("接受通道值", chalVal, "当前循环的次数", i)
			wg.Done()
		}(i)
	}

	fmt.Println(`main~`)
	close(ch)
	wg.Wait()
}
