package main

import (
	"bufio"
	"day9/chat/commons"
	"encoding/json"
	"fmt"
	"os"
)

func sendTextMessage(msgText string) (err error) {
	userSendMsg := commons.UserSendMessageReq{
		UserId: client.userId,
		Data:   msgText,
	}
	data, err := json.Marshal(userSendMsg)
	if err != nil {
		return
	}
	msg := commons.Message{
		Cmd:  commons.UserSendMessageReqCmd,
		Data: string(data),
	}
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	err = writePackage(data)
	return
}

func enterTalk() {
	var msgText string
	fmt.Println("input text")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		msgText = input.Text()
	}
	err := sendTextMessage(msgText)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func showMenu() {
	fmt.Println("0. exit")
	fmt.Println("1. List OnLine User")
	fmt.Println("2. Talk")
	var sel int
	fmt.Scanf("%d\n", &sel)
	switch sel {
	case 0:
		os.Exit(0)
	case 1:
		outputUserOnline()
	case 2:
		enterTalk()
	}
}

func logic() {
	showMenu()
}
