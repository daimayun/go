package function

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// HashHmac 使用HMAC方法生成键控哈希值[hash_hmac]
func HashHmac(algo, msg, key string) string {
	var m hash.Hash
	if algo == "md5" {
		m = hmac.New(md5.New, []byte(key))
	} else if algo == "sha1" {
		m = hmac.New(sha1.New, []byte(key))
	} else if algo == "sha256" {
		m = hmac.New(sha256.New, []byte(key))
	} else if algo == "sha512" {
		m = hmac.New(sha512.New, []byte(key))
	} else {
		m = hmac.New(sha256.New, []byte(key))
	}
	m.Write([]byte(msg))
	return hex.EncodeToString(m.Sum(nil))
}
