package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		trim := strings.Trim(readString, "\r\n")
		if trim == "exit" {
			fmt.Println("客户端推出")
			break
		}
		write, err := dial.Write([]byte(readString))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		fmt.Println(write)
	}
}
