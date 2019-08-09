package main

import (
	"fmt"
)

func swpe(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func swpe1(a int, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 100, 200
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	// swpe(&a, &b)
	// a, b = swpe1(a, b)
	a, b = b, a
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
