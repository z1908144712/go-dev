package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bigAdd(a, b string) (sum string) {
	len_a := len(a)
	len_b := len(b)
	if len_a == 0 && len_b == 0 {
		sum = "0"
		return
	}
	a_index := len_a - 1
	b_index := len_b - 1
	flag := 0
	for a_index >= 0 && b_index >= 0 {
		c1 := a[a_index] - '0'
		c2 := b[b_index] - '0'
		c3 := int(c1) + int(c2) + flag
		flag = c3 / 10
		sum = string(c3%10+'0') + sum
		a_index--
		b_index--
	}
	for a_index >= 0 {
		c1 := a[a_index] - '0'
		c3 := int(c1) + flag
		flag = c3 / 10
		sum = string(c3%10+'0') + sum
		a_index--
	}
	for b_index >= 0 {
		c2 := b[b_index] - '0'
		c3 := int(c2) + flag
		flag = c3 / 10
		sum = string(c3%10+'0') + sum
		b_index--
	}
	if flag == 1 {
		sum = "1" + sum
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	strs := strings.Split(string(result), "+")
	if len(strs) != 2 {
		fmt.Println("please input a+b")
		return
	}
	c1 := strings.TrimSpace(strs[0])
	c2 := strings.TrimSpace(strs[1])
	fmt.Println(bigAdd(c1, c2))
}
