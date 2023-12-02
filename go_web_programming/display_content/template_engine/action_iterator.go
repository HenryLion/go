package main

import (
	"net/http"
	"text/template"
)

func processiterator(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("iterator.html")
	daysOfWeek := []string{"Monday", "TuesDay", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", processiterator)
	server.ListenAndServe()
}
