package function

import (
	"strings"
	"unicode/utf8"
)

// StrLengthCutAndSplitJoint 字符串长度截取并拼接处理
func StrLengthCutAndSplitJoint(str string, cutLength int, splitJointStr ...string) string {
	if cutLength == 0 {
		return str
	}
	if utf8.RuneCountInString(str) <= cutLength {
		return str
	}
	strRune := []rune(str)
	s := string(strRune[0:cutLength])
	if len(splitJointStr) > 0 {
		s += splitJointStr[0]
	}
	return s
}

// CheckStringLength 判断字符串长度是否在规定范围内
func CheckStringLength(str string, length int) bool {
	if utf8.RuneCountInString(str) <= length {
		return true
	}
	return false
}

// ToBigCamelCase 字符串转大驼峰格式
func ToBigCamelCase(str string, signs ...string) (res string) {
	sign := "_"
	if len(signs) > 0 {
		sign = signs[0]
	}
	strArr := strings.Split(str, sign)
	for _, v := range strArr {
		res += strings.ToUpper(string(v[0])) + v[1:]
	}
	return
}

// ToSmallCamelCase 字符串转小驼峰格式
func ToSmallCamelCase(str string, signs ...string) (res string) {
	sign := "_"
	if len(signs) > 0 {
		sign = signs[0]
	}
	strArr := strings.Split(str, sign)
	for k, v := range strArr {
		if k == 0 {
			res += strings.ToLower(string(v[0])) + v[1:]
		} else {
			res += strings.ToUpper(string(v[0])) + v[1:]
		}
	}
	return
}
