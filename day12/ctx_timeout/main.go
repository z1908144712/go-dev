package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{
			r:   resp,
			err: err,
		}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("timeout")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Println(string(out))
	}
	return
}

func main() {
	process()
}
