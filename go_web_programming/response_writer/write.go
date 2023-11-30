package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
			<head>Go Web Programming</head>
			<body><h1>Hello World333</h1></body>
			</html>`
	//str2 := `{"name":"han"}`
	w.Write([]byte(str))
}

// WriteHeader 只用来设置http返回码
func writeHeadExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "No such service, try next door")
}

// 卧槽，这里访问http://127.0.0.1:8080/redirect 可以直接跳转到google页面
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com") // 用来设置header里的一些key:value
	w.WriteHeader(302)
	w.Write([]byte("nageren"))
}

func returnjson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("nagekey", "nagevalue")
	post := &Post{
		User:    "pencil han",
		Threads: []string{"first", "second"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeadExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/returnjson", returnjson)
	server.ListenAndServe()
}
