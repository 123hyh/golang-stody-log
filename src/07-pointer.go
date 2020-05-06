package main

/**
指针（指向内存的地址)
 */
func main()  {
	i:=17
	p:=&i
	//取值得指针
	println(p)
	//根据指针取值
	println(*p)
	newAddess()
}
//new 申请内存地址
func newAddess()  {
	a:=new(int)
	*a = 18
	println(a)
}