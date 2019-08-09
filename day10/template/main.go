package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	p := Person{
		Name: "person",
		Age:  18,
	}
	if err = t.Execute(w, p); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
