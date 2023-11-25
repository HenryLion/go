package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello1!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("check user infomation...")
		h(w, r)
	}
}

func main() {
	serve := http.Server{
		Addr: ":8080",
	}

	// 此处看似绑定的log函数，实则是绑定的log函数返回的那个函数，并且在返回的那个函数中调用了hello1函数
	http.HandleFunc("/hello", protect(log(hello1)))
	serve.ListenAndServe()
}
