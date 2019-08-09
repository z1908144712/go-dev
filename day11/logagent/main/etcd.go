package main

import (
	"context"
	"encoding/json"
	"go_dev/day11/logagent/commons"
	"go_dev/day11/logagent/tailf"
	"strings"
	"time"

	"go.etcd.io/etcd/mvcc/mvccpb"

	"github.com/astaxie/beego/logs"

	"go.etcd.io/etcd/clientv3"
)

var (
	etcdClient *commons.EtcdClient
)

func InitEtcd(addr, key string) (collectPaths []commons.CollectPath, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	etcdClient = &commons.EtcdClient{
		Client: cli,
	}
	if !strings.HasSuffix(key, "/") {
		key = key + "/"
	}
	for _, ip := range localIP {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		etcdKey := key + ip
		etcdClient.Keys = append(etcdClient.Keys, etcdKey)
		resp, err := cli.Get(ctx, etcdKey)
		cancel()
		if err != nil {
			logs.Error("get config from %s failed\n", etcdKey)
			continue
		}
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectPaths)
				if err != nil {
					logs.Error(err)
					continue
				}
			}
		}
	}

	InitEtcdWatcher()

	return
}

func InitEtcdWatcher() {
	for _, v := range etcdClient.Keys {
		go Watcher(v)
	}
}

func Watcher(key string) {
	for {
		rch := etcdClient.Client.Watch(context.Background(), key)
		var collectPaths []commons.CollectPath
		for wrch := range rch {
			getConfSuccess := true
			for _, ev := range wrch.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s]' config delete")
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &collectPaths)
					if err != nil {
						logs.Error("key[%s] Unmarshal failed,", err)
						getConfSuccess = false
						continue
					}
					logs.Debug("config update, key:%s, vaule:%s", ev.Kv.Key, ev.Kv.Value)
				}
			}
			if getConfSuccess {
				tailf.UpdateConf(collectPaths)
			}
		}

	}
}
