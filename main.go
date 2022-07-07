package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is Index Page!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// env GOOS=linux GOARCH=amd64 go build -o webserver
