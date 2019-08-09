package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			continue
		}
		fmt.Println(msg)
	}
}
