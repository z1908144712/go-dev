package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	var input int
	for {
		fmt.Scanf("%d\n", &input)
		switch {
		case input == n:
			fmt.Println("you are right!")
			return
		case input > n:
			fmt.Println("bigger!")
		case input < n:
			fmt.Println("less!")
		}
	}
}
