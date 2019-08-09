package main

import (
	"go_dev/day11/logagent/commons"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error(err)
			time.Sleep(time.Second * 2)
			continue
		}
	}
}

func sendToKafka(msg *commons.TextMsg) (err error) {
	kafka.SendToKafka(msg)
	return
}
