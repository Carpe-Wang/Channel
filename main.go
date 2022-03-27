package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	//阻塞等待用户连接

	for { //接受多个用户
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		////接受用户请求
		//buf := make([]byte, 1024) //设置1024缓冲区
		//n, err1 := conn.Read(buf)
		//if err1 != nil {
		//	fmt.Println(err1)
		//	continue
		//}
		//fmt.Printf("buf:", string(buf[:n])) //读了多少
		//defer conn.Close()

		go HandleConn(conn)
	}
	//接受用户请求
}

func HandleConn(conn net.Conn) {
	//获取网络信息
	add := conn.RemoteAddr().String()
	add = add
	fmt.Println("Connection is successful ")
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("has err:", err)
		return
	}
	fmt.Println(string(buf[:n]))

	//把数据转化为大写，并发送给用户
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}
