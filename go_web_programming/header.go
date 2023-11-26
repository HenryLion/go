package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/headers", header)
	http.HandleFunc("/body", body) // 此处虽然只是将body写入应答，但是go语言会自动写入一些header信息，所以看到的输出中多了一些header信息
	server.ListenAndServe()
}
