package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day11/logagent/commons"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	etcdClient *clientv3.Client
)

type LogInfo struct {
	AppId      int    `db:"app_id"`
	AppName    string `db:"app_name"`
	LogId      int    `db:"log_id"`
	LogPath    string `db:"log_path"`
	Topic      string `db:"topic"`
	CreateTime string `db:"create_time"`
	Status     int    `db:"status"`
}

func InitEtcd(client *clientv3.Client) {
	etcdClient = client
}

func GetAllLogInfo() (logList []LogInfo, err error) {
	err = Db.Select(&logList, "select a.app_id, b.app_name, a.log_id, a.log_path,a.topic,a.create_time, a.status from tbl_log_info a, tbl_app_info b where a.app_id=b.app_id")
	return
}

func CreateLog(info *LogInfo) (err error) {
	conn, err := Db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}
		conn.Commit()
	}()
	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil {
		return
	}
	if len(appId) == 0 {
		err = errors.New("app_id is nil")
		return
	}
	info.AppId = appId[0]
	_, err = conn.Exec("insert into tbl_log_info(app_id,log_path, topic)values(?,?,?)", info.AppId, info.LogPath, info.Topic)
	if err != nil {
		return
	}
	return
}

func SetEtcdConf(etcdKey string, info *LogInfo) (err error) {
	var collectPaths []commons.CollectPath
	collectPaths = append(collectPaths, commons.CollectPath{
		LogPath: info.LogPath,
		Topic:   info.Topic,
	})
	data, err := json.Marshal(collectPaths)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = etcdClient.Put(ctx, etcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdClient.Get(ctx, etcdKey)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range resp.Kvs {
		fmt.Println(k, v)
	}
	return
}
