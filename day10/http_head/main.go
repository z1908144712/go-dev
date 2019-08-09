package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := []string{
		"http://www.baidu.com",
		"http://googole.com",
		"http://taobao.com",
	}
	for _, v := range url {
		res, err := http.Head(v)
		if err != nil {
			fmt.Println(v, "\t", err)
			continue
		}
		fmt.Println(v, "\t", res.Status)
	}
}
