package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1210"
)

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye", sayBye)
	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 2,
		Handler:      mux,
	}
	log.Println("Starting httpserver at" + Addr)
	// 运行
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye, this is httpserver"))
}
