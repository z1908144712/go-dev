package main

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

func convertLogLevel(logLevel string) int {
	switch logLevel {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "error":
		return logs.LevelError
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["log_level"] = convertLogLevel(appConfig.LogLevel)
	configData, err := json.Marshal(config)
	if err != nil {
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configData))
	return
}
