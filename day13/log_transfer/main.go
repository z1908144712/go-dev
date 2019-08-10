package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	err := initConfig("./conf/log_transfer.conf", "ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = initLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = initKafka()
	if err != nil {
		logs.Error(err)
		return
	}
	err = initES()
	if err != nil {
		logs.Error(err)
		return
	}
	// err = run()
	// if err != nil {
	// 	logs.Error(err)
	// 	return
	// }
	logs.Warn("log_transfer is exited")
}
