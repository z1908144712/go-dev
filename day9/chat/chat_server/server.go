package main

import (
	"fmt"
	"net"
)

func runServer(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	client := &Client{
		conn: conn,
	}
	err := client.Process()
	if err != nil {
		fmt.Println("client:", err)
		return
	}
}
