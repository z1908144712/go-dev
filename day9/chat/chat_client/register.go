package main

import (
	"day9/chat/commons"
	"encoding/json"
	"fmt"
)

func register() (err error) {
	var msg commons.Message
	msg.Cmd = commons.UserLoginReqCmd
	var registerReq commons.RegisterReq
	user := commons.User{
		UserId: client.userId,
		Nick:   fmt.Sprintf("user%d", client.userId),
		Sex:    "1",
		Passwd: client.passwd,
		Status: commons.UserOffline,
	}
	registerReq.User = user
	data, err := json.Marshal(registerReq)
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
	if msg.Cmd != commons.UserRegisterResCmd {
		err = commons.ErrInvaildCmd
		return
	}
	var registerRes commons.RegisterRes
	err = json.Unmarshal([]byte(msg.Data), &registerRes)
	if err != nil {
		return
	}
	if registerRes.Code != 200 {
		err = fmt.Errorf(registerRes.Error)
		return
	}
	return
}
