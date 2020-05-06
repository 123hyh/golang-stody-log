package main

import (
	"fmt"
	"os"
)

type stutent struct {
	id    int
	name  string
	class string
}

func newStutent(id int, name, class string) *stutent {
	return &stutent{
		id:    id,
		name:  name,
		class: class,
	}
}

type stutentMrg struct {
	allStutent []*stutent
}

// 创建管理
func newStutentMrg() *stutentMrg {
	return &stutentMrg{
		allStutent: make([]*stutent, 0, 100),
	}
}

// 添加成员
func (s *stutentMrg) addStutent(newStu *stutent) {
	s.allStutent = append(s.allStutent, newStu)
}

//编辑成员
func (s *stutentMrg) editStutent(id int, name, class string) {
	for i, val := range s.allStutent {
		if val.id == id {
			s.allStutent[i] = &stutent{
				id:    id,
				name:  name,
				class: class,
			}
		}
	}
}

//展示所有
func (s *stutentMrg) showStutents() {
	for _, val := range s.allStutent {
		fmt.Printf("学号：%#d 姓名：%#s 班级：%#s\n", val.id, val.name, val.class)
	}
}
func (s *stutentMrg) exit() {
	println("退出")
	os.Exit(0)
}
