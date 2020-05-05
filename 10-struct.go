package main

import (
	"fmt"
	"unsafe"
)

/*
结构体
*/

//结构体实例化
type personType struct {
	name, city string
	age        int
}

func (p *personType) getName() {
	fmt.Printf("方法和接收者%v\n", p.name)
}
func (p *personType) setAge(newAge int) {
	p.age = newAge
	fmt.Printf("指针类型的接收者\n%v\n")
}

func main() {
	// 类型定义
	type myInt int

	//类型别名
	type myUint = uint
	var a myUint = 1
	fmt.Printf("%v\n", a)

	var person personType
	person.name = "HuangYuHui"
	person.city = "shenzhen"
	person.age = 18
	fmt.Printf("结构体实例化%#v\n", person)

	//	匿名结构体
	var tiger struct{ color, name string }
	tiger.name = "小金"
	tiger.color = "金"
	fmt.Printf("匿名结构体%#v\n", tiger)

	//指针结构体
	p2 := new(personType)
	p2.name = "黄"
	p2.age = 19
	p2.city = "shenzhen"
	fmt.Printf("指针结构体%#v\n", p2)

	//	使用键值对初始化
	p3 := personType{
		name: "黄裕辉",
		age:  18,
		city: "深圳",
	}
	fmt.Printf("价值对初始化%#v\n", p3)

	//	使用值的列表初始化
	p4 := &personType{
		"黄",
		"韶关",
		18,
	}
	fmt.Printf("使用值的列表初始化%v\n", p4)

	fmt.Printf("空结构体%v\n", unsafe.Sizeof(struct {
	}{}))

	//	构造函数
	construct := func(name, city string, age int) *personType {
		return &personType{
			name: name,
			city: city,
			age:  age,
		}
	}
	fmt.Printf("通过构造函数实例化结构体%v\n", construct("黄", "深圳", 18))

	//	map 结构体
	m1 := make(map[string]*personType)
	m1["黄裕辉"] = &personType{"黄裕辉", "shenzhen", 18}
	fmt.Printf("map 结构体%#v\n", m1["黄裕辉"])

	//	切片 结构体
	s1 := []*personType{}
	s1 = append(s1, &personType{
		"黄裕辉",
		"shenzhen",
		18,
	})
	fmt.Printf("切片 结构体%#v\n", s1[0])

	//方法和接收者
	newPerson := func(name, city string, age int) *personType {
		return &personType{
			name,
			city,
			age,
		}
	}
	ph1 := newPerson("黄裕辉", "韶关", 27)
	fmt.Printf("%v\n", ph1)
	ph1.getName()
	ph1.setAge(20)
	fmt.Printf("%#v\n", ph1)

	//	结构体的匿名字段
	type anony struct {
		string
		int
	}

	//	嵌套结构体
	type Address struct {
		province string
		city     string
	}
	type User struct {
		name, gender string
		Address
	}
	hyh := User{
		name:   "黄裕辉",
		gender: "男",
	}
	hyh.Address.province = "广东"
	hyh.Address.city = "韶关"
	fmt.Printf("嵌套结构体%#v\n", hyh)

	//	结构体的继承
	dog := Dog{
		feet:   18,
		Animal: &Animal{name: "小金"},
	}
	dog.move()
}

type Animal struct {
	name string
}
type Dog struct {
	feet int
	*Animal
}

func (a *Animal) move() {
	fmt.Printf("结构体的继承%v\n", a.name)
}
