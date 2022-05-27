package des

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// ECBEncrypt AES的ECB模式加密
func ECBEncrypt(data, key []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.ECBEncrypt(block, data, padding)
}

// ECBDecrypt AES的ECB模式解密
func ECBDecrypt(data, key []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.ECBDecrypt(block, data, padding)
}
