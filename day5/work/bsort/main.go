package main

import (
	"fmt"
)

func main() {
	a := [...]int{23, 43, 1, 23, 67, 88, 90, 34, 56, 37, 44, 26}
	bsort(a[:])
	fmt.Println(a)
}

func bsort(b []int) {
	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b)-i; j++ {
			if b[j-1] > b[j] {
				b[j-1], b[j] = b[j], b[j-1]
			}
		}
	}
}
