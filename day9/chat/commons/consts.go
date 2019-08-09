package commons

const (
	UserLoginReqCmd       = "user_login_req"
	UserLoginResCmd       = "user_login_res"
	UserRegisterReqCmd    = "user_register_req"
	UserRegisterResCmd    = "user_register_res"
	UserStatusNotifyCmd   = "user_status_notify"
	UserSendMessageReqCmd = "user_send_message_req"
	UserRecvMessageCmd    = "user_recv_message"
)

const (
	UserOffline = 0
	UserOnline  = 1
)
