package main

import (
	"fmt"
)

func process(str string) bool {
	t := []rune(str)
	for i := 0; i < len(t)/2; i++ {
		if (t[i] != t[len(t)-1-i]) {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}