package main

import (
	"fmt"
)

func main() {
	a := [...]int{23, 43, 1, 23, 67, 88, 90, 34, 56, 37, 44, 26}
	ssort(a[:])
	fmt.Println(a)
}

func ssort(b []int) {
	for i := 0; i < len(b); i++ {
		min := i
		for j := i; j < len(b); j++ {
			if b[min] > b[j] {
				min = j
			}
		}
		if min != i {
			b[min], b[i] = b[i], b[min]
		}
	}
}