package main

import (
	"fmt"
)

func main() {
	a := [...]int{23, 43, 1, 23, 67, 88, 90, 34, 56, 37, 44, 26}
	isort(a[:])
	fmt.Println(a)
}

func isort(b []int) {
	for i := 1; i < len(b); i++ {
		for j := i; j > 0; j-- {
			if b[j] >= b[j-1] {
				break
			}
			b[j], b[j-1] = b[j-1], b[j]
		}
	}
}