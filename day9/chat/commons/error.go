package commons

import (
	"errors"
)

var (
	ErrUserExist     = errors.New("user exist")
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvaildPasswd = errors.New("passwd or username not right")
	ErrInvaildParams = errors.New("invaild params")
	ErrInvaildCmd    = errors.New("invaild cmd")
)
