package commons

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginReq struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterReq struct {
	User User `json:"user"`
}

type LoginRes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

type RegisterRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}

type UserSendMessageReq struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

type UserRecvMessage struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}
