package main

import (
	"fmt"
	"net/http"
)

// 定义两个函数，其函数签名满足所示条件，即可将函数绑定到serve
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	serve := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	serve.ListenAndServe()
}
