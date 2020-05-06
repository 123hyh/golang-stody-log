package main

import "fmt"

func main() {
	zoons := CreateZoonManager()
	for {
		showMenu()
		fmt.Println("请输入您的操作")
		var input uint8
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			value := inputValue()
			zoons.AddZoon(value)
		case 2:
			zoons.ShowAllZoon()
		case 3:
			zoons.Exit()
		default:
			fmt.Println("没有该操作~")

		}
	}
}
func showMenu() {
	println("1、添加动物")
	println("2、展示动物")
	println("3、退出~")
}
func inputValue() *zoon {
	name := ""
	color := ""
	var sex uint8
	fmt.Println("请输入名字")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入颜色")
	fmt.Scanf("%s\n", &color)
	fmt.Println("请输入性别")
	fmt.Scanf("%d\n", &sex)
	var value = make(map[string]interface{})

	value["name"] = name
	value["color"] = color
	value["sex"] = sex
	return NewZoon(name,color,sex)
}
