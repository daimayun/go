package aes

import (
	"crypto/aes"
	"github.com/daimayun/go/openssl"
)

// CTREncrypt AES的CTR模式加密
func CTREncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CTREncrypt(block, data, iv, padding)
}

// CTRDecrypt AES的CTR模式解密
func CTRDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CTRDecrypt(block, data, iv, padding)
}
