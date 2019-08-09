package main

import (
	"day9/chat/commons"
	"fmt"
)

var onlineUsers = make(map[int]*commons.User)

func outputUserOnline() {
	if len(onlineUsers) == 0 {
		return
	}
	fmt.Println("Online User List:")
	for id, _ := range onlineUsers {
		if id == client.userId {
			continue
		}
		fmt.Println("user:", id)
	}
}

func updateOnlineUser(userStatusNotify commons.UserStatusNotify) {
	user, ok := onlineUsers[userStatusNotify.UserId]
	if !ok {
		user = &commons.User{}
		user.UserId = userStatusNotify.UserId
	}
	user.Status = userStatusNotify.Status
	onlineUsers[user.UserId] = user
}
