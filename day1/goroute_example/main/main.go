package main

import (
	"fmt"
	"day1/goroute_example/goroute"
)

func main() {
	pipe := make(chan int, 2)
	go goroute.Add(100, 200, pipe)
	sum := <- pipe
	sub := <- pipe
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}