package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World2!")
}

func main() {
	h := MyHandler{}
	serve := http.Server{
		Addr:    ":8080",
		Handler: &h,
	}
	serve.ListenAndServe()
}
