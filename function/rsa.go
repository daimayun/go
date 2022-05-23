package function

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
)

// RsaEncrypt 使用对方的公钥的数据,只有对方的私钥才能解开
func RsaEncrypt(plain string, publicKey string) (cipherByte []byte, err error) {
	var (
		pubKeyValue interface{}
		encryptOAEP []byte
	)
	msg := []byte(plain)
	// 解码公钥
	pubBlock, _ := pem.Decode([]byte(publicKey))
	// 读取公钥
	pubKeyValue, err = x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return
	}
	pub := pubKeyValue.(*rsa.PublicKey)
	// 加密数据方法:不用使用EncryptPKCS1v15方法加密,源码里面推荐使用EncryptOAEP,因此这里使用安全的方法加密
	encryptOAEP, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, msg, nil)
	if err != nil {
		return
	}
	cipherByte = encryptOAEP
	return
}

// RsaDecrypt 使用私钥解密公钥加密的数据
func RsaDecrypt(cipherByte []byte, privateKey string) (plainText string, err error) {
	var (
		priKey      *rsa.PrivateKey
		decryptOAEP []byte
	)
	// 解析出私钥
	priBlock, _ := pem.Decode([]byte(privateKey))
	priKey, err = x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		return
	}
	// 解密RSA-OAEP方式加密后的内容
	decryptOAEP, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherByte, nil)
	if err != nil {
		return
	}
	plainText = string(decryptOAEP)
	return
}
