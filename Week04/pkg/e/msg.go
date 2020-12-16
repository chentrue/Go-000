package e

var MsgFlags = map[int]string{
	SUCCESS : "ok",
	ERROR   : "fail",
	MAILBOX_CLOSED :"邮箱已关闭",
	MAILBOX_INVALID_RCPTTO : "收件人地址无效",
}

func GetMsg(code int) string  {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}