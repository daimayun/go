package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CBCEncrypt CBC模式加密
func CBCEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	// pkcs5填充
	data = openssl.Pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText
}

// CBCDecrypt CBC模式解密
func CBCDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	// pkcs5填充
	cryptText = openssl.Pkcs5UnPadding(cryptText)
	return cryptText
}
