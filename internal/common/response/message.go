package response

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[IntervalServer] = "服务器开小差啦,稍后再来试一试"
}

func GetMessage(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return message[IntervalServer]
	}
}
