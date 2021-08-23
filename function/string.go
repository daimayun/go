package function

import "unicode/utf8"

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
