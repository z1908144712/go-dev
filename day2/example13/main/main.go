package main

import (
	"fmt"
)

func main() {
	var m int32 = 34000
	var n int16
	n = int16(m)
	fmt.Println("m=", m)
	fmt.Println("n=", n)
}
