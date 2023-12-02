package main

import (
	"net/http"
	"text/template"
)

func processinclude(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "pencilhan")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", processinclude)
	server.ListenAndServe()
}
