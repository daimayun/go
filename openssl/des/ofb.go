package des

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// OFBEncrypt DES的OFB模式加密
func OFBEncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.OFBEncrypt(block, data, iv, padding)
}

// OFBDecrypt DES的OFB模式解密
func OFBDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.OFBDecrypt(block, data, iv, padding)
}
