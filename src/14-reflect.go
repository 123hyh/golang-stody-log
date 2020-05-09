package main

import (
	"fmt"
	"reflect"
)

/**
反射
*/

func reflectType(param interface{}) {
	//获取值的类型对象
	v := reflect.TypeOf(param)
	fmt.Printf("%v \n", v.Kind())
}

//原始值的值信息
func reflectValue(x interface{}) {

	v := reflect.ValueOf(x).Elem() //根据指针取对应的值 elem()

	t := v.Kind() // 获取值对于的类型种类
	switch t {
	case reflect.Float32:
		nVal := float32(v.Float()) // 强制转换类型
		fmt.Printf("%v\n", nVal)
		fmt.Printf("float32: %v %T\n", t, nVal)

	case reflect.Int8:
		nVal := int8(v.Int())
		fmt.Printf("int8: %v\n", nVal)
		//修改对应的值
		v.SetInt(100)

	case reflect.Bool:
		fmt.Println("是布尔", v.Bool())

	case reflect.Map:
		fmt.Printf("%v\n", v)
	}
}

func main() {
	var a float32 = 3.141592653
	reflectType(a)
	var b int8 = 123
	var c bool = false
	var m = make(map[string]int)
	m["name"] = 12

	// 传入指针
	reflectValue(&a)
	reflectValue(&b)
	reflectValue(&c)
	reflectValue(&m)
	//判断指针是否为空
	fmt.Println(reflect.ValueOf(&a).IsNil())
}
