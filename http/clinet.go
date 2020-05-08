package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	//创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, //谈话时间
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲时间
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls 握手时间
		ExpectContinueTimeout: 1 * time.Second,  //状态码超时时间
	}
	//创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, //最大请求超时,
		Transport: transport,
	}
	//请求数据
	get, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	//读取数据
	if err != nil {
		panic(err)
	}
	fmt.Println(string((all)))

}
