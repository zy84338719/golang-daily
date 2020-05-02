package main

import (
	"fmt"
	"net"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Println("server %s", conn.RemoteAddr().String())
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client exit")
			return
		}
		fmt.Println(string(buf[:read]))
		sprint := fmt.Sprint(string(buf[:read]))
		if strings.Trim(sprint, "\r\n") == "server exit" {
			return
		}
	}
}

func main() {
	fmt.Println("服务器监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("-----")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Println("Accept() con=%v ip=%v", accept, accept.RemoteAddr().String())
		}
		go process(accept)
	}
}
