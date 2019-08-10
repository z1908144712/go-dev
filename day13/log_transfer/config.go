package main

import (
	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr string
	Topic     string
	ESAddr    string
	LogPath   string
	LogLevel  string
}

var (
	appConfig *LogConfig
)

func initConfig(filename, confType string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		return
	}
	appConfig = &LogConfig{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}
	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs/log_transfer.log"
	}
	appConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.KafkaAddr) == 0 {
		appConfig.KafkaAddr = "127.0.0.1:9092"
	}
	appConfig.Topic = conf.String("kafka::topic")
	if len(appConfig.Topic) == 0 {
		appConfig.Topic = "nginx_log"
	}
	appConfig.ESAddr = conf.String("es::server_addr")
	if len(appConfig.ESAddr) == 0 {
		appConfig.ESAddr = "127.0.0.1:9200"
	}
	return
}
