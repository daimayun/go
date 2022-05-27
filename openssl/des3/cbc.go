package des3

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CBCEncrypt DES3的CBC模式加密
func CBCEncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CBCEncrypt(block, data, iv, padding)
}

// CBCDecrypt DES3的CBC模式解密
func CBCDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CBCDecrypt(block, data, iv, padding)
}
