package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Studen struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}
type Class struct {
	Title   string    `json:"title"`
	Count   int       `json:"count"`
	Studens []*Studen `json:"studens"`
}

func (c *Class) setStudens(list []*Studen) {
	c.Studens = make([]*Studen, len(list))
	copy(c.Studens, list)
}

//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

func main() {

	newClass := &Class{Title: "九年级", Count: 20, Studens: make([]*Studen, 0, 20)}
	mapSlice := func(name string) []*Studen {
		list := make([]*Studen, 0)
		for i := 0; i < 20; i++ {
			data := &Studen{Age: i, Name: name + strconv.Itoa(i), Gender: i % 2}
			list = append(list, data)
		}
		return list
	}
	newClass.setStudens(mapSlice("mff"))
	for _, val := range newClass.Studens {
	println(val)
	}
	// json 序列化
	data, _ := json.Marshal(newClass)
	// json 反序列化
	n1 := &Class{}
	json.Unmarshal([]byte(data), n1)
	fmt.Printf("%s\n", data)
}
