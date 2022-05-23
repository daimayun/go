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
