package main

import (
	"fmt"
	"net/http"
)

// 可以直接返回HTML浏览器就可以展示
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<html><body><h1>Hello, World! %s </h1></body></html>", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
