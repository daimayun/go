package des3

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CTREncrypt DES3的CTR模式加密
func CTREncrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CTREncrypt(block, data, iv, padding)
}

// CTRDecrypt DES3的CTR模式解密
func CTRDecrypt(data, key, iv []byte, padding openssl.PaddingType) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return openssl.CTRDecrypt(block, data, iv, padding)
}
