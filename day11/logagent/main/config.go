package main

import (
	"errors"
	"go_dev/day11/logagent/commons"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *commons.Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var collectPath commons.CollectPath
	collectPath.LogPath = conf.String("collect::log_path")
	if len(collectPath.LogPath) == 0 {
		err = errors.New("invaild collect::log_path")
		return
	}
	collectPath.Topic = conf.String("collect::topic")
	if len(collectPath.Topic) == 0 {
		err = errors.New("invaild collect::topic")
		return
	}
	appConfig.CollectPaths = append(appConfig.CollectPaths, collectPath)
	return
}

func loadConf(filename, confType string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		return
	}
	appConfig = &commons.Config{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}
	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs/logagent.log"
	}
	appConfig.ChanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.ChanSize = 100
	}
	appConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.KafkaAddr) == 0 {
		appConfig.KafkaAddr = "127.0.0.1:9092"
	}
	appConfig.EtcdAddr = conf.String("etcd::addr")
	if len(appConfig.EtcdAddr) == 0 {
		appConfig.EtcdAddr = "127.0.0.1:2379"
	}
	appConfig.EtcdKey = conf.String("etcd::confKey")
	if len(appConfig.EtcdKey) == 0 {
		appConfig.EtcdKey = "/logagent/conf/"
	}
	err = loadCollectConf(conf)
	return
}
