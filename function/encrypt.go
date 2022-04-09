package function

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5 md5加密
func Md5(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha1 sha1加密
func Sha1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
