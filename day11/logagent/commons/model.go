package commons

import (
	"sync"

	"github.com/hpcloud/tail"
	"go.etcd.io/etcd/clientv3"
)

type Config struct {
	LogLevel     string
	LogPath      string
	ChanSize     int
	KafkaAddr    string
	EtcdAddr     string
	EtcdKey      string
	CollectPaths []CollectPath
}

type CollectPath struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}

type TailObj struct {
	Tail     *tail.Tail
	Conf     CollectPath
	ExitChan chan int
}

type TailObjMgr struct {
	Tails   []*TailObj
	MsgChan chan *TextMsg
	Lock    sync.Mutex
}

type TextMsg struct {
	Text  string
	Topic string
}

type EtcdClient struct {
	Client *clientv3.Client
	Keys   []string
}
