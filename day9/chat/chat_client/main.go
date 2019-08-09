package main

import (
	"day9/chat/commons"
	"fmt"
	"net"
	"os"
)

type Client struct {
	userId   int
	passwd   string
	conn     net.Conn
	recvMsgs chan commons.UserRecvMessage
}

var client *Client

func init() {
	client = &Client{}
	client.recvMsgs = make(chan commons.UserRecvMessage, 1000)
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("client connect err", err)
		os.Exit(-1)
	}
	client.conn = conn
}

func main() {
	fmt.Scanf("%d %s\n", &client.userId, &client.passwd)
	err := login()
	if err != nil {
		fmt.Println("login failed,", err)
		if err.Error() == commons.ErrUserNotExist.Error() {
			err = register()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("user register success")
		}
		return
	}
	fmt.Println("login success")
	outputUserOnline()
	go processServerMessage()
	for {
		logic()
	}
	defer client.conn.Close()
}
