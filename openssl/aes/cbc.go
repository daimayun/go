package aes

import (
	"crypto/aes"
	"github.com/daimayun/go/openssl"
)

// CBCEncrypt AES的CBC模式加密
func CBCEncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CBCEncrypt(block, data, iv, padding)
}

// CBCDecrypt AES的CBC模式解密
func CBCDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CBCDecrypt(block, data, iv, padding)
}
