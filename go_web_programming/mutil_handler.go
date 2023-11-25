package main

import (
	"fmt"
	"net/http"
)

// 定义两个结构体，且结构体实现ServeHTTP函数，然后这结构体就可以作为handler绑定到serve上了
type HelloHandler struct{}
type WorldHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: ":8080",
	}

	// 此处是给serve绑定了两个struct对象，对struct的要求是，其有一个ServeHTTP函数
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
