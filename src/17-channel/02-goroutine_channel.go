package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//生成 0~10的数(限制单向通道)
func f1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

//计算它的平方(限制单向通道)
func f2(ch1 <-chan int, ch2 chan<- int) {
	// 从通道中取值
	for ret := range ch1 {
		ch2 <- (ret * ret)
	}
	close(ch2)
	wg.Done()
}
func main() {
	wg.Add(2)
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)

	for ret := range ch1 {
		fmt.Println(ret)
	}
	wg.Wait()
}
/**
通道通信
*/
/*func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	//赋值
	go func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}(ch1)
	//取值
	go func(ch <-chan int, ch1 chan<- int) {
		for i := range ch {
			ch1 <- i * i
		}
		close(ch1)
	}(ch1, ch2)
	for i := range ch2 {
		fmt.Println(i)
	}
}
*/