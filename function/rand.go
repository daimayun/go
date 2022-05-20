package function

import (
	cr "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strconv"
)

// RandString 32位随机字符串 [不推荐使用]
func RandString() string {
	rand.Seed(now.UnixNano())
	return Md5(fmt.Sprintf("%s%s", strconv.FormatInt(now.UnixNano(), 10), strconv.Itoa(rand.Intn(1000000))))
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
	return Md5(strconv.FormatInt(now.UnixNano(), 10) + base64.URLEncoding.EncodeToString(b) + s)
}
