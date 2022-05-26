package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// OFBEncrypt OFB模式加密
func OFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	// pkcs5填充
	data = openssl.Pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))
	blockMode := cipher.NewOFB(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// OFBDecrypt OFB模式解密
func OFBDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewOFB(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)
	// pkcs5填充
	cryptText = openssl.Pkcs5UnPadding(cryptText)
	return cryptText
}
