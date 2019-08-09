package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("err:", err)
	}
	wc, sc, nc, oc := count(string(str))
	fmt.Printf("wordCount:%d\nspaceCount:%d\nnumberCount:%d\notherCount:%d\n", wc, sc, nc, oc)
}
