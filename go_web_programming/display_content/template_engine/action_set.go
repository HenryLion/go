package main

import (
	"net/http"
	"text/template"
)

func process_set(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("set.html")
	t.Execute(w, "pencilhan")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process_set)
	server.ListenAndServe()
}
