package function

import (
	cr "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// RandString 32位随机字符串 [不推荐使用]
func RandString() string {
	nowUnixNano := time.Now().UnixNano()
	rand.Seed(nowUnixNano)
	return Md5(fmt.Sprintf("%s%s", strconv.FormatInt(nowUnixNano, 10), strconv.Itoa(rand.Intn(1000000))))
}

// UniqueId 生成Guid字串
func UniqueId(v ...interface{}) string {
	s := ""
	if len(v) > 0 {
		s = fmt.Sprintf("%s", v[0])
	}
	b := make([]byte, 48)
	if _, err := io.ReadFull(cr.Reader, b); err != nil {
		return RandString()
	}
	return Md5(strconv.FormatInt(time.Now().UnixNano(), 10) + base64.URLEncoding.EncodeToString(b) + s)
}

// RandShortStringByDate 根据当前日期生成6位短字符串[以秒为单位不重复]
func RandShortStringByDate() (str string) {
	slice := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "J", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	now := time.Now()
	date := now.Format(TimeLayoutYMDHIS)[2:]
	j := 0
	for i := 1; i <= 6; i++ {
		s := date[j:(j + 2)]
		si, _ := strconv.Atoi(s)
		if si > 61 {
			str += s
		} else {
			str += slice[si]
		}
		j += 2
	}
	return
}

// RandFixedString 根据[0-9]数字生成一位随机固定字符串[没有大小写oO]
func RandFixedString(i int) string {
	if !(i >= 0 && i <= 9) {
		return ""
	}
	slice := [][]string{
		{"0", "a", "k", "v", "F", "Q"},
		{"1", "b", "l", "w", "G", "R"},
		{"2", "c", "m", "x", "H", "S"},
		{"3", "d", "n", "y", "I", "T"},
		{"4", "e", "p", "z", "J", "U"},
		{"5", "f", "q", "A", "K", "V"},
		{"6", "g", "r", "B", "L", "W"},
		{"7", "h", "s", "C", "M", "X"},
		{"8", "i", "t", "D", "N", "Y"},
		{"9", "j", "u", "E", "P", "Z"},
	}
	gRand := rand.New(rand.NewSource(time.Now().UnixNano()).(rand.Source64))
	SliceShuffle(&slice[i])
	return slice[i][gRand.Intn(6)]
}

// RandNum 返回随机数字字符
func RandNum(length int) (str string) {
	for i := 0; i < length; i++ {
		str += strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano() + int64(i))).Intn(10))
	}
	return
}
