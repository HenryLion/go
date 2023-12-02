package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func processaction(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("condition.html")
	rand.Seed(time.Now().Unix())
	num := rand.Intn(10)
	t.Execute(w, num > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", processaction)
	server.ListenAndServe()
}
