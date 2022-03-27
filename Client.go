package main

import (
	"fmt"
	"net"
)

func main() {
	//自动连接服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	read := make([]byte, 1024)
	read = append(read, 'A', 'r', 'e', ' ', 'y', 'o', 'u', ' ', 'o', 'k')
	conn.Write(read)
	defer conn.Close()
	demo(conn)

}

func demo(conn net.Conn) {
	buf := make([]byte, 2048)
	n, _ := conn.Write(buf)
	conn.Read(buf[:n])
	fmt.Println(buf)
}
