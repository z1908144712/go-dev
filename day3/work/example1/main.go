package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	for i := n; i < m; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
		}
	}
}
