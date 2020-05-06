package main

import "fmt"

type DogClass struct {
}

func (d *DogClass) call() {
	fmt.Println("汪汪汪~")
}
func (d *DogClass) move() {
	fmt.Println(`狗在移动`)
}

type Cat struct {
	name string
}

func (c *Cat) call() {
	fmt.Println("喵喵喵~")
}
func (c *Cat) move() {
	fmt.Println(`猫在移动`)
}

type hit interface {
	call()
}

// 接口嵌套
type animal interface {
	hit
}

func call(arg hit) {
	arg.call()
}

func main() {

	var d1 animal
	d2 := &DogClass{}
	c2 := &Cat{name: "小金"}
	// 值类型 和指针接受者
	d1 = c2

	call(d1)
	call(d2)
	call(c2)
	nullInterface()
}

//空接口
func nullInterface() {
	interParam(1, "o", true)
	var i interface{}
	i = true
	fmt.Println(i)

	//	空接口类型作为  map value的类型
	m1 := make(map[string]interface{})
	m1["name"] = "hyh"
	m1["age"] = 18
	m1["hobby"] = []string{"羽毛球"}
	fmt.Printf("%#v\n", m1)

	//	类型断言
	ret, ok := i.(string)
	if ok {
		fmt.Printf("%v\n", ret)
	}
	//	使用switch 断言
	switch v := i.(type) {
	case string:
		fmt.Printf("字符串类型 value: %v\n", v)
	case bool:
		fmt.Printf("布尔类型 value: %v\n", v)
	default:
		fmt.Println(`猜不到`)
	}
}

// 空接口类型作为 函数参数
func interParam(param ...interface{}) {
	println(`123`)
}
