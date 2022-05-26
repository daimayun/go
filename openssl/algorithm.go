package openssl

import "bytes"

// Pkcs5Padding pkcs5补码算法
func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	return append(ciphertext, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

// Pkcs5UnPadding pkcs5减码算法
func Pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	return origData[:(length - int(origData[length-1]))]
}

// ZeroPadding zero补码算法
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	return append(ciphertext, bytes.Repeat([]byte{0}, padding)...)
}

// ZeroUnPadding zero减码算法
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
