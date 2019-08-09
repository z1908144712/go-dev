package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "logcollect.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Port:", port)
	log_level, err := conf.Int("log::log_level")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Log_level:", log_level)
	log_path := conf.String("log::log_path")
	fmt.Println("Log_path:", log_path)
}
