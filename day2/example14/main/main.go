package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float32())
	}
}
