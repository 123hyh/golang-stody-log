package main

import "fmt"

func showMenu() {
	println(`1、添加学员`)
	println(`2、编辑学员`)
	println(`3、展示学员`)
	println(`4、退出学员`)
}

func getInput() *stutent {
	var (
		id    int
		name  string
		class string
	)
	println("输入id")
	fmt.Scanf("%d\n", &id)
	println("输入姓名")
	fmt.Scanf("%s\n", &name)
	println("输入班级")
	fmt.Scanf("%s\n", &class)
	return newStutent(id, name, class)
}
func main() {
	mrg := newStutentMrg()
	for {
		showMenu()
		var input int
		fmt.Printf("输入的操作")
		fmt.Scanf("%d\n", &input)

		switch input {
		case 1:
			inputStutent := getInput()
			mrg.addStutent(inputStutent)
			fmt.Printf("新增成功：%s\n", mrg.allStutent)
		case 2:
			inputStutent := getInput()
			mrg.editStutent(inputStutent.id, inputStutent.name, inputStutent.class)
			mrg.showStutents()
		case 3:
			mrg.showStutents()
		case 4:
			mrg.exit()
		default:
			fmt.Printf("没有你想要的操作\n")
		}
	}
}
