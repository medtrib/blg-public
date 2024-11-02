package constants

// StatusMessage 状态和提示信息的结构体
type StatusMessage struct {
	Code    int32
	Message string
}

// 定义常量来表示 StatusCodes 的键
const (
	StatusSuccess             = "Success"
	StatusError               = "Error"
	StatusSnBlank             = "SnBlank"
	StatusCodeInvalid         = "CodeInvalid"
	StatusAppIdInvalid        = "AppIdInvalid"
	StatusDeviceInfoErr       = "DeviceInfoError"
	StatusOpenIdInvalid       = "OpenIdInvalid"
	StatusPhoneInvalid        = "PhoneInvalid"
	StatusUuIDInvalid         = "UuidInvalid"
	StatusSignNameInvalid     = "SignNameInvalid"
	StatusAccessKeyInvalid    = "AccessKey"
	StatusAccessSecretInvalid = "AccessSecret"
	StatusSmsCodeInvalid      = "SmsCodeInvalid"
)

// StatusCodes 定义状态和提示信息
var StatusCodes = map[string]StatusMessage{
	"Success": {
		Code:    1,
		Message: "操作成功",
	},
	"Error": {
		Code:    0,
		Message: "操作失败",
	},
	"SnBlank": {
		Code:    0,
		Message: "sn 不能为空",
	},
	"CodeInvalid": {
		Code:    0,
		Message: "Code无效的代码",
	},
	"AppIdInvalid": {
		Code:    0,
		Message: "无效的 AppId",
	},
	"DeviceInfoError": {
		Code:    0,
		Message: "设备暂未开通",
	},
	"OpenIdInvalid": {
		Code:    0,
		Message: "无效的 OpenId",
	},
	"UuiDInvalid": {
		Code:    0,
		Message: "无效的 UuId",
	},
	"SignNameInvalid": {
		Code:    0,
		Message: "无效的 SignName",
	},
	"SmsCodeInvalid": {
		Code:    0,
		Message: "无效的 SmsCode",
	},
}
