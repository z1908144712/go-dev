package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006/1/2 15:4:05"))
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Println("cost", (end-start)/1000)
}
