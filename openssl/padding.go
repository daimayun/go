package openssl

import (
	"bytes"
	"errors"
)

// PaddingType 填充类型
type PaddingType string

var (
	PKCS5 PaddingType = "PKCS5"
	PKCS7 PaddingType = "PKCS7"
	Zero  PaddingType = "Zero"
)

var unPaddingError = errors.New("UnPadding error")

func Padding(padding PaddingType, src []byte, blockSize int) []byte {
	switch padding {
	case PKCS5:
		src = PKCS5Padding(src, blockSize)
	case PKCS7:
		src = PKCS7Padding(src, blockSize)
	case Zero:
		src = ZerosPadding(src, blockSize)
	}
	return src
}

func UnPadding(padding PaddingType, src []byte) ([]byte, error) {
	switch padding {
	case PKCS5:
		return PKCS5UnPadding(src)
	case PKCS7:
		return PKCS7UnPadding(src)
	case Zero:
		return ZerosUnPadding(src)
	}
	return src, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	return PKCS7Padding(src, blockSize)
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	return PKCS7UnPadding(src)
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return src, unPaddingError
	}
	unPadding := int(src[length-1])
	if length < unPadding {
		return src, unPaddingError
	}
	return src[:(length - unPadding)], nil
}

func ZerosPadding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func ZerosUnPadding(src []byte) ([]byte, error) {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1], nil
		}
	}
}
