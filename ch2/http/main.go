package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter Sample")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
