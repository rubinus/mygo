package main

import (
	"io"
	"mygo/morerequest/do"
	"net/http"

	"github.com/json-iterator/go"
)

func main() {
	http.HandleFunc("/", One)
	//http.HandleFunc("/more", morequrest)
	http.ListenAndServe("8080", nil)
}

func One(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello go </h1")
}

func morequrest(w http.ResponseWriter, r *http.Request) {
	result := do.DoWork(1)
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	be, _ := jsonIterator.Marshal(result)
	io.WriteString(w, string(be))
}
