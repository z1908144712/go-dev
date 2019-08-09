package main

import (
	"fmt"
)

func Sum(n int) int64 {
	var s int64 = 1
	var sum int64 = 0
	for i := 1; i <= n; i++ {
		s *= int64(i)
		fmt.Printf("%d! = %v\n", i, s)
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(Sum(n))
}
