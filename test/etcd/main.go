package etcd

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"test/etcd/discovery"
)

type service struct {
}

func main() {

	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器

	pb.RegisterMailServiceServer(s, &service{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。

	//etcd服务注册
	reg, err := discovery.NewService(discovery.ServiceInfo{
		Name: "g.srv.mail",
		IP:   "127.0.0.1:8972", //grpc服务节点ip
	}, []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"}) // etcd的节点ip
	if err != nil {
		log.Fatal(err)
	}
	go reg.Start()

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
