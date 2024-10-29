package utils

import (
	"encoding/json"
)

// ToJsonString 通用转换Json字符串函数
func ToJsonString(message interface{}) (string, error) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
