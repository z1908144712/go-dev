package main

import (
	"net"

	"github.com/astaxie/beego/logs"
)

var (
	localIP []string
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logs.Error(err)
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4 != nil {
				localIP = append(localIP, ipnet.IP.String())
			}
		}
	}
}
