package function

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"io/ioutil"
	"os"
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

// HashFile 文件哈希[优先使用]
func HashFile(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(buf)
	return hex.EncodeToString(h[:]), nil
}

// HashFileToMd5 文件哈希[MD5]
func HashFileToMd5(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := md5.Sum(buf)
	return hex.EncodeToString(h[:]), nil
}

// hashFile 文件哈希[不推荐使用]
func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hob := sha256.New()
	_, err = io.Copy(hob, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hob.Sum(nil)), nil
}
