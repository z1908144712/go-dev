package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello</h1>")
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}
