package main

import (
	"day9/chat/commons"
	"encoding/json"
	"fmt"
	"os"
)

func processServerMessage() {
	for {
		msg, err := readPackage()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		switch msg.Cmd {
		case commons.UserStatusNotifyCmd:
			var userStatusNotify commons.UserStatusNotify
			err = json.Unmarshal([]byte(msg.Data), &userStatusNotify)
			if err != nil {
				fmt.Println(err)
				return
			}
			updateOnlineUser(userStatusNotify)
			outputUserOnline()
		case commons.UserRecvMessageCmd:
			var userRecvMessage commons.UserRecvMessage
			err = json.Unmarshal([]byte(msg.Data), &userRecvMessage)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(userRecvMessage.UserId, ":", userRecvMessage.Data)
		}
	}
}
