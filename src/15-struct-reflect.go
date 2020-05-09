package main

import (
	"fmt"
	"reflect"
)

type stutent struct {
	Name  string `json:"s_name" ini:"i_name"`
	Score int    `json:"s_score" ini:"i_score"`
}

func (s stutent) Call() {
	fmt.Println(s.Name + `叫了~`)
}

func main() {
	stu1 := stutent{
		Name:  "hyh",
		Score: 85,
	}
	fmt.Println(`------------------根据索引获取字段信息-----------------------`)

	//通过反射获取结构体的所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name: %v; kind: %v\n", t.Name(), t.Kind())

	//	遍历结构体遍历的所有字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("name: %v; tag: %v; type: %v\n", f.Name, f.Tag, f.Type)

		//	取 json 和 ini tag
		fmt.Println(`获取json和ini tag名称：`, f.Tag.Get("json"), f.Tag.Get("ini"))
	}

	fmt.Println(`------------------根据名称取结构体字段信息-----------------------`)
	//	根据名字取结构体字段信息
	nt, ok := t.FieldByName("Name")
	if ok {
		fmt.Println(nt.Type, nt.Name)
	}
	printMethod(stu1)
}
func printMethod(s interface{}) {
	fmt.Println(`------------------根据反射获取结构体中的方法的函数-----------------------`)
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name, t.Method(i).Type)
		//调用其方法
		args := []reflect.Value{}
		v.Method(i).Call(args)
	}
}
