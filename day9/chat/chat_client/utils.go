package main

import (
	"day9/chat/commons"
	"encoding/binary"
	"encoding/json"
	"errors"
)

func readPackage() (msg commons.Message, err error) {
	var buf [8192]byte
	n, err := client.conn.Read(buf[:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	if err != nil {
		return
	}
	packageLen := binary.BigEndian.Uint32(buf[:4])
	n, err = client.conn.Read(buf[:packageLen])
	if n != int(packageLen) {
		err = errors.New("read body failed")
		return
	}
	if err != nil {
		return
	}
	err = json.Unmarshal(buf[:packageLen], &msg)
	if err != nil {
		return
	}
	return
}

func writePackage(data []byte) (err error) {
	packageLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], packageLen)
	n, err := client.conn.Write(buf[:])
	if n != 4 {
		err = errors.New("write header failed")
		return
	}
	if err != nil {
		return
	}
	n, err = client.conn.Write(data)
	if err != nil {
		return
	}
	if n != int(packageLen) {
		err = errors.New("write body failed")
		return
	}
	return
}
