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

// RandString 32位随机字符串
func RandString() string {
	rand.Seed(time.Now().UnixNano())
	return Md5(fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UnixNano(), 10), strconv.Itoa(rand.Intn(1000000))))
}

// UniqueId 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(cr.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
