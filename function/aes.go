package function

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// 只支持16、24、32位，分别对应AES-128，AES-192，AES-256 加密方法
var password = []byte("774D58AB5192D2556F5C1D39C6E049E5")

// PKCS7Padding PKCS7 填充模式[AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个]
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unPadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unPadding)], nil
	}
}

// AesEncrypt AES加密
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

// AesDecrypt AES解密
func AesDecrypt(encrypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(encrypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, encrypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// AesEncryptToBase64 加密后为Base64格式的字符串
func AesEncryptToBase64(str string, key ...string) (string, error) {
	if len(key) > 0 {
		password = []byte(key[0])
	}
	result, err := AesEncrypt([]byte(str), password)
	if err != nil {
		return "", err
	}
	return Base64Encode(string(result)), err
}

// AesDecryptByBase64 对Base64格式的字符串解密
func AesDecryptByBase64(str string, key ...string) (string, error) {
	if len(key) > 0 {
		password = []byte(key[0])
	}
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	var b []byte
	//执行AES解密
	b, err = AesDecrypt(pwdByte, password)
	if err == nil {
		return string(b), nil
	}
	return "", err
}
