package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// CFBEncrypt CFB模式加密
func CFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	// pkcs5填充
	data = openssl.Pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))
	blockMode := cipher.NewCFBDecrypter(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// CFBDecrypt CFB模式解密
func CFBDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCFBEncrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)
	// pkcs5填充
	cryptText = openssl.Pkcs5UnPadding(cryptText)
	return cryptText
}
