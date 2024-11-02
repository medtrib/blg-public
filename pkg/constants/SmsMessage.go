package constants

const (
	// SmsRedisSendMessage 短信发送Key
	SmsRedisSendMessage = "SMS_Redis_SEND_MESSAGE:"
)

// GetSmsKey 返回缓存值数据
func GetSmsKey(smsKey, phone string) string {
	return smsKey + phone
}

// SmsType 定义短信类型
var SmsType = map[string]StatusMessage{
	"Success": {
		Code:    1,
		Message: "操作成功",
	},
	"Error": {
		Code:    0,
		Message: "操作失败",
	},
}
