package main

import (
	"fmt"
	"io"
	"net/http"
)

func process1(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Println(r.MultipartForm.File["uploaded"])
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process1)
	server.ListenAndServe()
}
