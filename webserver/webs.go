package main

import (
	"fmt"
	"io"

	jsoniter "github.com/json-iterator/go"

	"mygo/morerequest/do"
	"net/http"
)

func main() {
	http.HandleFunc("/", One)
	//http.HandleFunc("/more", morequrest)

	err := http.ListenAndServe(":8080", nil)

	fmt.Print(222)

	if err != nil {
		panic(err)
	}
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
