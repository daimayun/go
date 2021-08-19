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
