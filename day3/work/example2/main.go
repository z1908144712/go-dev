package main

import (
	"fmt"
)

func isNumber(n int) bool {
	i := n % 10
	j := (n / 10) % 10
	k := (n / 100) % 10
	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}

func main() {
	var n, m int
	fmt.Scanf("%d-%d", &n, &m)
	for i := n; i <= m; i++ {
		if isNumber(i) {
			fmt.Println(i)
		}
	}
}
