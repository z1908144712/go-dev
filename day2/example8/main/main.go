package main

import (
	"fmt"
)

func main() {
	a := 5
	b := make(chan int, 3)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}