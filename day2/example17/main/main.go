package main

import (
	"fmt"
)

func reverse(str *string) {
	var tmp byte
	strs := []byte(*str)
	n := len(strs)
	for i := 0; i < n/2; i++ {
		tmp = strs[i]
		strs[i] = strs[n-i-1]
		strs[n-i-1] = tmp
	}
	*str = string(strs)
}

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)
	reverse(&str3)
	fmt.Println(str3)
}
