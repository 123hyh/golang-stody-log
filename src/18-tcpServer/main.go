package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 处理完之后要关闭这个连接
	
}
func main() {
	// 1 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//等待客户端建立连接
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue

		}
		//开启一个单独的 goroutine处理连接
		process(connect)
	}
}
