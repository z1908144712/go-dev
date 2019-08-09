package main

import (
	"day9/chat/commons"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	id   int
	buf  [8192]byte
}

func (p *Client) readPackage() (msg commons.Message, err error) {
	n, err := p.conn.Read(p.buf[:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	if err != nil {
		return
	}
	packageLen := binary.BigEndian.Uint32(p.buf[:4])
	n, err = p.conn.Read(p.buf[:packageLen])
	if n != int(packageLen) {
		err = errors.New("read body failed")
		return
	}
	if err != nil {
		return
	}
	err = json.Unmarshal(p.buf[:packageLen], &msg)
	if err != nil {
		return
	}
	return
}

func (p *Client) writePackage(data []byte) (err error) {
	packageLen := uint32(len(data))
	binary.BigEndian.PutUint32(p.buf[:4], packageLen)
	n, err := p.conn.Write(p.buf[:4])
	if n != 4 {
		err = errors.New("write header failed")
		return
	}
	if err != nil {
		return
	}
	n, err = p.conn.Write(data)
	if err != nil {
		return
	}
	if n != int(packageLen) {
		err = errors.New("write body failed")
		return
	}
	return
}

func (p *Client) Process() error {
	for {
		msg, err := p.readPackage()
		if err != nil {
			clientMgr.DelCilent(p.id)
			return err
		}
		err = p.processMsg(msg)
		if err != nil {
			continue
		}
	}
}

func (p *Client) NotifyUserMessage(userSendMessageReq commons.UserSendMessageReq) {
	var reMsg commons.Message
	reMsg.Cmd = commons.UserRecvMessageCmd
	var notify commons.UserRecvMessage
	notify.UserId = userSendMessageReq.UserId
	notify.Data = userSendMessageReq.Data
	data, err := json.Marshal(notify)
	if err != nil {
		return
	}
	reMsg.Data = string(data)
	data, err = json.Marshal(reMsg)
	if err != nil {
		return
	}
	err = p.writePackage(data)
	if err != nil {
		return
	}
	return
}

func (p *Client) processUserSendMessage(msg commons.Message) (err error) {
	var userSendMessageReq commons.UserSendMessageReq
	err = json.Unmarshal([]byte(msg.Data), &userSendMessageReq)
	if err != nil {
		return
	}
	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userSendMessageReq.UserId {
			continue
		}
		client.NotifyUserMessage(userSendMessageReq)
	}
	return
}

func (p *Client) processMsg(msg commons.Message) (err error) {
	switch msg.Cmd {
	case commons.UserLoginReqCmd:
		err = p.login(msg)
	case commons.UserRegisterReqCmd:
		err = p.register(msg)
	case commons.UserSendMessageReqCmd:
		err = p.processUserSendMessage(msg)
	default:
		err = errors.New("unsupport message cmd")
		return
	}
	return
}

func (p *Client) login(msg commons.Message) (err error) {
	defer func() {
		var reMsg commons.Message
		reMsg.Cmd = commons.UserLoginResCmd
		var loginRes commons.LoginRes
		if err != nil {
			loginRes.Code = 500
			loginRes.Error = fmt.Sprintf("%v", err)
		} else {
			loginRes.Code = 200
			userMap := clientMgr.GetAllUsers()
			for userId, _ := range userMap {
				loginRes.User = append(loginRes.User, userId)
			}
		}
		data, err := json.Marshal(loginRes)
		if err != nil {
			return
		}
		reMsg.Data = string(data)
		data, err = json.Marshal(reMsg)
		if err != nil {
			return
		}
		err = p.writePackage(data)
		if err != nil {
			return
		}
		p.NotifyOthersUserOnline(p.id)
		return
	}()
	var loginReq commons.LoginReq
	err = json.Unmarshal([]byte(msg.Data), &loginReq)
	if err != nil {
		return
	}
	_, err = mgr.Login(loginReq.Id, loginReq.Passwd)
	if err != nil {
		return
	}
	p.id = loginReq.Id
	clientMgr.AddCilent(loginReq.Id, p)
	return
}

func (p *Client) register(msg commons.Message) (err error) {
	defer func() {
		var reMsg commons.Message
		reMsg.Cmd = commons.UserRegisterResCmd
		var registerRes commons.RegisterRes
		registerRes.Code = 200
		if err != nil {
			registerRes.Code = 500
			registerRes.Error = fmt.Sprintf("%v", err)
		}
		data, err := json.Marshal(registerRes)
		if err != nil {
			return
		}
		reMsg.Data = string(data)
		data, err = json.Marshal(reMsg)
		if err != nil {
			return
		}
		err = p.writePackage(data)
		if err != nil {
			return
		}
		return
	}()
	if msg.Cmd != commons.UserRegisterReqCmd {
		err = commons.ErrInvaildCmd
		return
	}
	var registerReq commons.RegisterReq
	err = json.Unmarshal([]byte(msg.Data), &registerReq)
	if err != nil {
		return
	}
	err = mgr.Register(&registerReq.User)
	if err != nil {
		return
	}
	return
}

func (p *Client) NotifyOthersUserOnline(userId int) {
	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userId {
			continue
		}
		client.NotifyUserOnline(userId)
	}
}

func (p *Client) NotifyUserOnline(userId int) {
	var reMsg commons.Message
	reMsg.Cmd = commons.UserStatusNotifyCmd
	var notify commons.UserStatusNotify
	notify.UserId = userId
	notify.Status = commons.UserOnline
	data, err := json.Marshal(notify)
	if err != nil {
		return
	}
	reMsg.Data = string(data)
	data, err = json.Marshal(reMsg)
	if err != nil {
		return
	}
	err = p.writePackage(data)
	if err != nil {
		return
	}
	return
}
