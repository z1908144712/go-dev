package main

import (
	"fmt"
	"go_dev/day11/logagent/kafka"
	"go_dev/day11/logagent/tailf"

	"github.com/astaxie/beego/logs"
)

func main() {
	err := loadConf("./conf/logagent.conf", "ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = initLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	appConfig.CollectPaths, err = InitEtcd(appConfig.EtcdAddr, appConfig.EtcdKey)
	if err != nil {
		logs.Error(err)
		return
	}
	err = tailf.InitTail(appConfig.CollectPaths, appConfig.ChanSize)
	if err != nil {
		logs.Error(err)
		return
	}
	err = kafka.InitKafka(appConfig.KafkaAddr)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info("initialize success")
	err = serverRun()
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Info("program exit")
}
