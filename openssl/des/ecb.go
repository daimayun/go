package des

import (
	"crypto/des"
	"github.com/daimayun/go/openssl"
)

// ECBEncrypt ECB模式加密
func ECBEncrypt(data, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	bs := block.BlockSize()
	data = openssl.Pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

// ECBDecrypt ECB模式解密
func ECBDecrypt(data, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	// pkcs5填充
	out = openssl.Pkcs5UnPadding(out)
	return out
}
