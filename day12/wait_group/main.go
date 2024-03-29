package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(wg, i)
	}
	wg.Wait()
	fmt.Println("done")
}

func calc(wg *sync.WaitGroup, i int) {
	fmt.Println(i)
	time.Sleep(time.Second)
	wg.Done()
}
