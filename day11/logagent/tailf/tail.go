package tailf

import (
	"go_dev/day11/logagent/commons"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

var (
	tailObjMgr *commons.TailObjMgr
)

func UpdateConf(collectPaths []commons.CollectPath) {
	tailObjMgr.Lock.Lock()
	defer tailObjMgr.Lock.Unlock()
	for _, collectPath := range collectPaths {
		isRunning := false
		for _, tail := range tailObjMgr.Tails {
			if collectPath.LogPath == tail.Conf.LogPath && collectPath.Topic == tail.Conf.Topic {
				isRunning = true
				break
			}
		}
		if isRunning {
			continue
		}
		err := NewTailTask(collectPath)
		if err != nil {
			logs.Error(err)
			continue
		}
	}
	var tailObjs []*commons.TailObj
	for _, tail := range tailObjMgr.Tails {
		status := false
		for _, collectPath := range collectPaths {
			if collectPath.LogPath == tail.Conf.LogPath && collectPath.Topic == tail.Conf.Topic {
				status = true
				break
			}
		}
		if !status {
			tail.ExitChan <- 1
		}
		tailObjs = append(tailObjs, tail)
	}
	tailObjMgr.Tails = tailObjs
}

func NewTailTask(collectPath commons.CollectPath) (err error) {
	tail, tailErr := tail.TailFile(collectPath.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if tailErr != nil {
		err = tailErr
		return
	}
	obj := &commons.TailObj{
		Tail:     tail,
		Conf:     collectPath,
		ExitChan: make(chan int, 1),
	}
	tailObjMgr.Tails = append(tailObjMgr.Tails, obj)
	go readFromTail(obj)
	return
}

func InitTail(collectPaths []commons.CollectPath, chanSize int) (err error) {
	tailObjMgr = &commons.TailObjMgr{
		MsgChan: make(chan *commons.TextMsg, chanSize),
	}
	if len(collectPaths) == 0 {
		logs.Error("invaild collect config")
		return
	}
	for _, v := range collectPaths {
		err = NewTailTask(v)
		if err != nil {
			logs.Error(err)
		}
	}
	return
}

func readFromTail(tailObj *commons.TailObj) {
	for {
		select {
		case line, ok := <-tailObj.Tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s", tailObj.Tail.Filename)
				time.Sleep(time.Second * 2)
				continue
			}
			textMsg := &commons.TextMsg{
				Text:  line.Text,
				Topic: tailObj.Conf.Topic,
			}
			tailObjMgr.MsgChan <- textMsg
		case <-tailObj.ExitChan:
			logs.Warn("tail will stop, conf:%v", tailObj.Conf)
			return
		}

	}
}

func GetOneLine() (msg *commons.TextMsg) {
	msg = <-tailObjMgr.MsgChan
	return
}
