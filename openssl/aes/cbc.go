package aes

import (
	"crypto/aes"
	"github.com/daimayun/go/openssl"
)

// CBCEncrypt CBC模式的加密
func CBCEncrypt(src, key, iv []byte, padding openssl.PaddingType) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
}

// CBCDecrypt CBC模式的解密
func CBCDecrypt(src, key, iv []byte, padding openssl.PaddingType) {
	//
}
