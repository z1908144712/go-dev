package main

import (
	"time"
)

func main () {
	for i := 0; i < 100; i++ {
		go test_route(i)
	}
	time.Sleep(time.Second)
}