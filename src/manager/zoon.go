package main

import (
	"fmt"
	"os"
)

type zoon struct {
	name  string
	color string
	sex   uint8
}

// 创建动物构造函数
func NewZoon(name, color string, sex uint8) *zoon {
	return &zoon{
		name:  name,
		color: color,
		sex:   sex,
	}
}

type zoonManeger struct {
	allZoon []*zoon
}

// 管理动物构造函数
func CreateZoonManager() *zoonManeger {
	return &zoonManeger{
		allZoon: make([]*zoon, 0, 100),
	}
}

// 添加动物
func (z *zoonManeger) AddZoon(v *zoon) {
	z.allZoon = append(z.allZoon, &zoon{name: v.name, color: v.color, sex: v.sex})
}

//展示所有动物
func (z *zoonManeger) ShowAllZoon() {
	for _, val := range z.allZoon {
		fmt.Printf("名字：%s 颜色 %s 性别 %d\n", val.name, val.color, val.sex)
	}
}

func (z *zoonManeger) Exit() {
	os.Exit(0)
}
