package main

import (
	"fmt"
	"time"
)

const (
	Man = 1
	Female = 2
)

func main() {
	var s int64
	for {
		s = time.Now().Unix()
		if (s % Female == 0) {
			fmt.Println("female")
		} else {
			fmt.Println("man")
		}
		time.Sleep(time.Second)
	}
}