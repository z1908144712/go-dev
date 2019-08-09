package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_dev/day11/logagent/commons"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	EtcdKey = "/logagent/conf/172.21.4.178"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()
	var collectPaths []commons.CollectPath
	collectPaths = append(collectPaths, commons.CollectPath{
		LogPath: "./logs/logagent.log",
		Topic:   "nginx_log",
	})
	collectPaths = append(collectPaths, commons.CollectPath{
		LogPath: "error.log",
		Topic:   "nginx_err",
	})
	data, err := json.Marshal(collectPaths)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range resp.Kvs {
		fmt.Println(k, v)
	}
}
