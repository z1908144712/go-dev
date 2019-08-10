package main

import (
	"go_dev/day13/web_admin/model"
	_ "go_dev/day13/web_admin/router"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.etcd.io/etcd/clientv3"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		return
	}
	model.InitDb(database)
	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	model.InitEtcd(cli)
	return
}

func main() {
	err := initDb()
	if err != nil {
		logs.Error(err)
	}
	err = initEtcd()
	if err != nil {
		logs.Error(err)
	}
	beego.Run()
}
