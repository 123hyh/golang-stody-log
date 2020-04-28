package main

import (
	"fmt"
	"strings"
)

func main() {
	baseType()
}
func baseType() {
	stringType()
	initUintType()

	/**
	* bool 类型 (默认为 false)
	 */
	var b1 bool = true
	fmt.Printf("bool类型：%v\n", b1)

}

/**
字符串类型
*/
func stringType() {
	fmt.Printf("start string：%v\n", "-----------------------------")
	/**
	字符串类型 (使用双引号包裹或者 反引号``)
	*/
	var s1 string = "hello"
	fmt.Printf("字符串类型：%#v\n", s1)

	// 转义字符
	var s2 string = `D:\codeStore\go-study\02-type.go`
	fmt.Printf("转义字符：%#v\n", s2)

	// 多行字符串
	var s3 string = `<div>
		<h1>多行字符串</h1>
	</div>`
	fmt.Printf("%#v\n", s3)

	/**
	字符串 api
	*/

	//字符串长度
	fmt.Printf("字符串长度：%v\n", len(s1))

	//	切割字符串
	var s4 string = "1,2,3,4,5"
	fmt.Printf("切割字符串：%T\n", strings.Split(s4, ","))

	//	包含某个字符
	fmt.Printf("是否包含某个字符：%v\n", strings.Contains(s4, "1"))

	//	前缀/后缀判断 字符
	fmt.Printf("判断前缀%v\n", strings.HasPrefix(s4, "1"))
	fmt.Printf("判断后缀%v\n", strings.HasSuffix(s4, "5"))

	//	查找某个字符串的索引
	fmt.Printf("查找首次出现某个字符串的索引:%v\n", strings.Index(s4, "5"))
	fmt.Printf("查找最后出现某个字符串的索引%v\n", strings.LastIndex(s4, ","))
	//	拼接
	var s5 []string = strings.Split(s4, ",")
	fmt.Printf("切割后拼接：%v\n", strings.Join(s5, ","))

	fmt.Printf("string end %v\n", "-----------------------------")
}

/**
int uint flot 类型
*/
func initUintType() {

	/**
	* int类型 (有字符)
	* 8, 16, 32, 64
	 */
	var n1 int = -1
	fmt.Printf("int类型：%v\n", n1)

	/**
	* uint类型 (无字符)
	* 8, 16, 32, 64
	 */
	var n2 uint = 1
	fmt.Printf("uint类型：%v\n", n2)

	/**
	flot 类型 ( 32, 64)
	*/
	var n3 float32 = 12.2
	var n4 float64 = 12.1215678678
	var n5 = float64(12.1215678678)
	fmt.Printf("flot32类型：%v\n", n3)
	fmt.Printf("flot64类型：%v\n", n4+n5)
}
