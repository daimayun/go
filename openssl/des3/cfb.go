package des3

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CFBEncrypt DES3的CFB模式加密
func CFBEncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CFBEncrypt(block, data, iv, padding)
}

// CFBDecrypt DES3的CFB模式解密
func CFBDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CFBDecrypt(block, data, iv, padding)
}
