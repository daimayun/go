package function

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

// UcFirst 仅开头字母大写[将字符串中的第一个字母转换成大写]
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 仅开头字母小写[将字符串中的第一个字母转换成小写]
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// UcWords 所有首字母大写[将字符串中每个单词的首字母转换成大写]
func UcWords(str string) (res string) {
	arr := strings.Split(str, " ")
	sign := ""
	for _, v := range arr {
		res += sign + UcFirst(v)
		sign = " "
	}
	return
}

// LcWords 所有首字母小写[将字符串中每个单词的首字母转换成小写]
func LcWords(str string) (res string) {
	arr := strings.Split(str, " ")
	sign := ""
	for _, v := range arr {
		res += sign + LcFirst(v)
		sign = " "
	}
	return
}

// Int64LengthPadding 数字长度不够左侧填补0并返回字符串
func Int64LengthPadding(i64 int64, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"d", i64)
}

// StringLengthPadding 字符串长度不够左侧填补0
func StringLengthPadding(str string, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"s", str)
}
