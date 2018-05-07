package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", One)
	http.ListenAndServe("8081", nil)
}

func One(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello go </h1")
}
