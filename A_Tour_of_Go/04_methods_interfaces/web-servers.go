package main

import (
	"fmt"
	"log"
	"net/http"
)

//
// Web servers
// Package http serves HTTP requests using any value that implements http.Handler:
//
// package http
//
// type Handler interface {
//     ServeHTTP(w ResponseWriter, r *Request)
// }
// In this example, the type Hello implements http.Handler.
//
// Visit http://localhost:4000/ to see the greeting.
//
// Note: This example won't run through the web-based tour user interface.
// To try writing web servers you may want to Install Go.
//
// Web 服务器
// 包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：
//
// package http
//
// type Handler interface {
//     ServeHTTP(w ResponseWriter, r *Request)
// }
// 在这个例子中，类型 Hello 实现了 http.Handler。
//
// 访问 http://localhost:4000/ 会看到来自程序的问候。
//
// *注意：* 这个例子无法在基于 web 的指南用户界面运行。为了尝试编写 web 服务器，可能需要安装 Go。
//

// 定义一个空类型，并实现http.ListenAndServe的Handler接口
type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
