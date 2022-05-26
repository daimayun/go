package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CTREncrypt CTR模式加密
func CTREncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	// pkcs5填充
	data = openssl.Pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))
	blockMode := cipher.NewCTR(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// CTRDecrypt CTR模式解密
func CTRDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCTR(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)
	// pkcs5填充
	cryptText = openssl.Pkcs5UnPadding(cryptText)
	return cryptText
}
