[TOC]

## 1. 变量声明
```go
//  var 
var s1 string = "hello word"

//const 
const (
	x = 1
	y = iota
)
```
[参考例子](./01-variable.go "例子")

------------


## 2. 数据类型
   -  ### 字符串类型
```go
// 必须是 双引号 or 反引号 
var s1 string = "hello word"
var s2 = "hello word"
// 多行字符串
var s3 string = `<div>
	<h1>hello word</h1>
</div>`
// 转义字符可以用 反引号
var s4 string = `D:\codeStore\golang-stoty.log`
```

------------


- ### 数字类型
-  int( 8, 16, 32 ,64 ) 无符号整形

```go
var n1 int = 18
var n2 = int8(18)
var n3 int16 = (18)
···
```

- uint( 8, 16, 32, 64 ) 有符号整形

```go
var n1 uint = 18
var n2 uint8 = (18)
var n3 uint16 = (18)
···
```
------------


- ### 浮点型
	- float( 32, 64)
	```go
	var n1 float32 = 12.12
	var n2  = float32(12.12)
	var n2  float64 = (12.121212312312)
	```

------------


- ### 其他数字类型
	- byte
	- rune
	- int
	- uint
	- uintptr

------------


- ###  布尔类型
```go
// 默认值 为 false
var b0 bool
var b1 bool = true
```

------------


