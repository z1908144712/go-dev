package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type AppInfo struct {
	AppId       int    `db:"app_id"`
	AppName     string `db:"app_name"`
	AppType     string `db:"app_type"`
	CreateTime  string `db:"create_time"`
	DevelopPath string `db:"develop_path"`
	IP          []string
}

var (
	Db *sqlx.DB
)

func InitDb(db *sqlx.DB) {
	Db = db
}

func GetAllAppInfo() (appList []AppInfo, err error) {
	err = Db.Select(&appList, "select app_id, app_name, app_type, create_time, develop_path from tbl_app_info")
	return
}

func GetIPInfoByName(appName string) (iplist []string, err error) {
	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", appName)
	if err != nil {
		return
	}
	if len(appId) == 0 {
		err = errors.New("select ip from tbl_app_info is nil")
		return
	}
	err = Db.Select(&iplist, "select ip from tbl_app_ip where app_id=?", appId[0])
	if err != nil {
		return
	}
	return
}

func CreateApp(info *AppInfo) (err error) {
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
	r, err := conn.Exec("insert into tbl_app_info(app_name, app_type, develop_path)values(?,?,?)", info.AppName, info.AppType, info.DevelopPath)
	if err != nil {
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		return
	}
	for _, ip := range info.IP {
		_, err = conn.Exec("insert into tbl_app_ip(app_id, ip)values(?,?)", id, ip)
		if err != nil {
			return
		}
	}
	return
}
