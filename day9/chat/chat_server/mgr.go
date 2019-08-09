package main

import (
	"day9/chat/commons"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	UserTable = "users"
)

type UserMgr struct {
	pool *redis.Pool
}

func NewUserMgr(pool *redis.Pool) *UserMgr {
	return &UserMgr{
		pool: pool,
	}
}

func (p *UserMgr) getUser(conn redis.Conn, id int) (user *commons.User, err error) {
	result, err := redis.String(conn.Do("HGet", UserTable, fmt.Sprintf("%d", id)))
	if err != nil {
		if err.Error() == redis.ErrNil.Error() {
			err = commons.ErrUserNotExist
		}
		return
	}
	user = &commons.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		return
	}
	return
}

func (p *UserMgr) Login(id int, passwd string) (user *commons.User, err error) {
	conn := p.pool.Get()
	defer conn.Close()
	user, err = p.getUser(conn, id)
	if err != nil {
		return
	}
	if user.UserId != id || user.Passwd != passwd {
		err = commons.ErrInvaildPasswd
		return
	}
	user.Status = commons.UserOnline
	user.LastLogin = fmt.Sprintf("%v", time.Now())
	fmt.Println("user", user.UserId, "login success!")
	return
}

func (p *UserMgr) Register(user *commons.User) (err error) {
	conn := p.pool.Get()
	defer conn.Close()
	if user == nil {
		err = commons.ErrInvaildParams
		return
	}
	_, err = p.getUser(conn, user.UserId)
	if err == nil {
		err = commons.ErrUserExist
		return
	}
	if err.Error() != commons.ErrUserNotExist.Error() {
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", UserTable, fmt.Sprintf("%d", user.UserId), data)
	if err != nil {
		return
	}
	fmt.Println("user", user.UserId, "register success!")
	return
}
