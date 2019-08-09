package main

import (
	"day9/chat/commons"
	"encoding/json"
	"fmt"
)

func login() (err error) {
	var msg commons.Message
	msg.Cmd = commons.UserLoginReqCmd
	var loginReq commons.LoginReq
	loginReq.Id = client.userId
	loginReq.Passwd = client.passwd
	data, err := json.Marshal(loginReq)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	err = writePackage(data)
	if err != nil {
		return
	}
	msg, err = readPackage()
	if err != nil {
		return
	}
	if msg.Cmd != commons.UserLoginResCmd {
		err = commons.ErrInvaildCmd
		return
	}
	var loginRes commons.LoginRes
	err = json.Unmarshal([]byte(msg.Data), &loginRes)
	if err != nil {
		return
	}
	if loginRes.Code != 200 {
		err = fmt.Errorf(loginRes.Error)
		return
	}
	for _, v := range loginRes.User {
		if v == client.userId {
			continue
		}
		user := &commons.User{
			UserId: v,
		}
		onlineUsers[v] = user
	}
	return
}
