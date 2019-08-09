package main

import (
	"fmt"
)

func calc(dataChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range dataChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}
	}
	exitChan <- true
}

func main() {
	dataChan := make(chan int, 10)
	resultChan := make(chan int, 10)
	exitChan := make(chan bool, 10)
	go func() {
		for i := 0; i < 100000; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()
	for i := 0; i < 10; i++ {
		go calc(dataChan, resultChan, exitChan)
	}
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}
		close(resultChan)
	}()
	for v := range resultChan {
		fmt.Println(v)
	}
}
