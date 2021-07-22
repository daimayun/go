package function

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5(str string) string {
	o := md5.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}

func Sha1(str string) string {
	o := sha1.New()
	o.Write([]byte(str))
	return hex.EncodeToString(o.Sum(nil))
}
