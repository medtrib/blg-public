package stringHelper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// IsPalindrome 判断字符串是否为回文
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	reversed := Reverse(s)
	return s == reversed
}

// CountOccurrences 统计子字符串在字符串中出现的次数
func CountOccurrences(s, substr string) int {
	return strings.Count(s, substr)
}

// GetSubstring 获取字符串的子字符串
func GetSubstring(s string, start, end int) string {
	if start < 0 || end > len(s) || start > end {
		return ""
	}
	return s[start:end]
}

// IndexOf 查找子字符串在字符串中的索引
func IndexOf(s, substr string) int {
	return strings.Index(s, substr)
}

// LastIndexOf 查找子字符串最后一次出现的索引
func LastIndexOf(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// ToTitle 将字符串转换为标题格式
func ToTitle(s string) string {
	return strings.Title(s)
}

// Reverse 反转字符串
func Reverse(s string) string {
	r := []rune(s) // 处理可能的Unicode字符
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// NormalizeWhitespace 将多个空白字符替换为单个空格
func NormalizeWhitespace(s string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(s), " ")
}

// IsStrongPassword 验证密码的强度
func IsStrongPassword(password string) bool {
	// 确保密码长度至少为8
	if len(password) < 8 {
		return false
	}

	// 使用正则表达式检查各类字符是否存在
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)                              // 检查大写字母
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)                              // 检查小写字母
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)                              // 检查数字
	hasSpecial := regexp.MustCompile(`[!@#%^&*()_+=\-{}[\]:;"'<>,.?/~]`).MatchString(password) // 检查特殊字符

	// 返回密码是否包含所有强度要求
	return hasUpper && hasLower && hasDigit && hasSpecial
}

// Trim 去掉字符串两端的空白字符
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToUpper 将字符串转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 将字符串转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Repeat 重复字符串指定次数
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// Split 将字符串按分隔符分割为切片
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join 将字符串切片按分隔符连接为一个字符串
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Replace 替换字符串中的子字符串
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// IsJSON 判断字符串是否为有效的 JSON 格式
func IsJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsPhoneNumber 验证字符串是否为有效的电话号码
func IsPhoneNumber(s string) bool {
	return regexp.MustCompile(`^\+?[0-9\s-]{7,15}$`).MatchString(s) // 简单的电话号码格式
}

// IsCreditCard 验证字符串是否为有效的信用卡号码
func IsCreditCard(s string) bool {
	return regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|6(?:011|5[0-9]{2})[0-9]{12}|3(?:0[0-5]|[68][0-9])[0-9]{11}|7[0-9]{15}|(?:2131|1800|35\d{3})\d{11})$`).MatchString(s)
}

// IsDate 验证字符串是否符合 YYYY-MM-DD 格式
func IsDate(s string) bool {
	return regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString(s)
}

// IsHex 判断字符串是否为有效的十六进制字符串
func IsHex(s string) bool {
	return regexp.MustCompile(`^0x?[0-9a-fA-F]+$`).MatchString(s)
}

// Contains 判断字符串是否包含某个子字符串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// StartsWith 判断字符串是否以某个子字符串开头
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith 判断字符串是否以某个子字符串结尾
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// IsAlpha 判断字符串是否仅包含字母
func IsAlpha(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)
}

// IsNumeric 判断字符串是否仅包含数字
func IsNumeric(s string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(s)
}

// IsAlphanumeric 判断字符串是否仅包含字母和数字
func IsAlphanumeric(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(s)
}

// IsEmail 验证字符串是否为有效的电子邮件格式
func IsEmail(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(s)
}

// IsURL 验证字符串是否为有效的URL格式
func IsURL(s string) bool {
	return regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`).MatchString(s)
}

// IsLength 验证字符串长度是否在指定范围内
func IsLength(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}

// IsBlank 判断字符串是否为空或仅包含空格
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank 判断字符串是否不为空且不只包含空格
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsEmpty 判断字符串是否为""（空字符串）
func IsEmpty(s string) bool {
	return s == ""
}

// IsNotEmpty 判断字符串是否不为""（非空字符串）
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// RandString 生成指定长度的随机字符串（仅包含大写字母）
func RandString(length int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := r.Intn(26) + 65 // 生成ASCII码范围为65-90的随机数（大写字母A-Z）
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// Md5Sum 获取字符串的MD5值
func Md5Sum(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GenRandomString 生成随机字符串
// length: 生成字符串的长度
// specialChar: 是否包含特殊字符（"yes" 或 "no"）
func GenRandomString(length int, specialChar string) string {
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "!@#%$*.="

	if specialChar == "yes" {
		letterBytes += special // 如果包含特殊字符，则添加到字符集中
	}

	chars := []byte(letterBytes)

	if length == 0 {
		return ""
	}

	clen := len(chars)
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // 用于存储随机字节
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return ""
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // 跳过此数字以避免模糊偏差
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// CryptPassword 加密密码
// password: 需要加密的密码
// salt: 加盐
func CryptPassword(password, salt string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", dk) // 返回十六进制字符串
}

// Base64Encode 简化Base64编码
func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data) // 使用URL安全的Base64编码
}

// Base64DecodeStripped 解码Base64字符串，处理缺失的填充字符
func Base64DecodeStripped(s string) (string, error) {
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i) // 添加缺失的填充字符
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	return string(decoded), err
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	// 应用id 生成雪花随机数
	node, _ := snowflake.NewNode(1)
	// 生成雪花id
	snowflakeId := node.Generate().String()
	return snowflakeId
}
