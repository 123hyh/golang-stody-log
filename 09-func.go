package main

import "fmt"

func main() {

	//类型简写
	typeLogogram := func(x, y int) {
		fmt.Printf("类型简写 %v, %v\n", x, y)
	}
	typeLogogram(1, 2)

	//可选参数
	choosableParam := func(x int, y ...int) {
		fmt.Printf("可选参数 %v\n", y)
	}
	choosableParam(1, 2, 3, 4, 5)

	//	多返回值
	moreRet := func(x, y int) (sum int, sub int) {
		sum = x + y
		sub = x - y
		return
	}
	sum, sub := moreRet(2, 1)
	fmt.Printf("多返回值%v %v\n", sum, sub)

	//	函数类型
	type calcType func(int, int) int
	var calcFun calcType = func(i int, i2 int) int {
		fmt.Printf("定义函数类型%v\n", 1)
		return 1
	}
	calcFun(1, 2)

	//	函数作为参数
	funcParam := func(x, y int, cb func(int, int)) {
		cb(x, y)
	}
	funcParam(1, 2, func(x int, y int) {
		fmt.Printf("函数作为参数%v %v\n", x, y)
	})

	//	函数作为返回值
	retFunc := func(x, y int) func() int {
		return func() int {
			fmt.Printf("函数作为返回值 %v %v\n", x, y)
			return x + y
		}
	}
	retFunc2 := retFunc(1, 2)
	retFunc2()

	//	自执行函数
	c := func(x, y int) int {

		fmt.Printf("自执行函数%v\n", x)
		return x + y
	}(1, 2)
	fmt.Printf("%v\n", c)
	//	闭包
	adder := func() func() {
		x := 0
		return func() {
			x++
			fmt.Printf("闭包%v\n\n", x)
		}
	}
	bdd := adder()
	bdd()
	bdd()

	//	defer 函数 （defer 关键字 逆序执行）
	testDefer := func() {
		fmt.Printf("同步%v\n", 1)
		defer fmt.Printf("异步%v\n\n", 2)
		fmt.Printf("同步%v\n", 3)
		defer fmt.Printf("异步%v\n", 4)
		defer fmt.Printf("异步%v\n", 5)
	}
	testDefer()
	f0 := func() (x int) {
		defer func() {
			x++
		}()
		return 6
	}
	fmt.Printf("%v\n\n", f0())
	testError()
	relation()

}

//	panic/recover 捕获错误
func testError() {
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func relation() {
	obj := make(map[string]int)
	conf := map[int]int{int('E'): 1, int('I'): 2, int('O'): 3, int('U'): 4}
	for key, val := range conf {
		conf[key+32] = val
	}
	names := []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	for _, key := range names {
		handler([]rune(key), obj, conf, key)
	}

	remain := func() int {
		count := 0
		for _, value := range obj {
			count += value
		}
		return 50 - count
	}()
	fmt.Printf("%v %v \n", remain, obj)
}
func handler(name []rune, obj map[string]int, conf map[int]int, user string) {
	for _, val := range name {
		if value, ok := conf[int(val)]; ok {
			if _, ok := obj[user]; ok {
				obj[user] += value
			} else {
				obj[user] = value
			}
		}
	}
}
