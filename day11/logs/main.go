package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	config["level"] = logs.LevelDebug
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Info("info")
	logs.Debug("debug")
	logs.Trace("trace")
	logs.Warn("warn")
}
