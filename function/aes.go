package function

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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

// AesEncryptByCBC AES加密[CBC模式]
func AesEncryptByCBC(origData []byte, key []byte) ([]byte, error) {
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

// AesDecryptByCBC AES解密[CBC模式]
func AesDecryptByCBC(encrypted []byte, key []byte) ([]byte, error) {
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

// AesEncryptByECB AES加密[ECB模式]
func AesEncryptByECB(origData []byte, key []byte) (encrypted []byte) {
	ciphers, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, ciphers.BlockSize(); bs <= len(origData); bs, be = bs+ciphers.BlockSize(), be+ciphers.BlockSize() {
		ciphers.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return
}

// AesDecryptByECB AES解密[ECB模式]
func AesDecryptByECB(encrypted []byte, key []byte) (decrypted []byte) {
	ciphers, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, ciphers.BlockSize(); bs < len(encrypted); bs, be = bs+ciphers.BlockSize(), be+ciphers.BlockSize() {
		ciphers.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

//// AesEncryptToBase64 加密后为Base64格式的字符串
//func AesEncryptToBase64(str string, key ...string) (string, error) {
//	if len(key) > 0 {
//		password = []byte(key[0])
//	}
//	result, err := AesEncrypt([]byte(str), password)
//	if err != nil {
//		return "", err
//	}
//	return Base64Encode(string(result)), err
//}
//
//// AesDecryptByBase64 对Base64格式的字符串解密
//func AesDecryptByBase64(str string, key ...string) (string, error) {
//	if len(key) > 0 {
//		password = []byte(key[0])
//	}
//	//解密base64字符串
//	pwdByte, err := base64.StdEncoding.DecodeString(str)
//	if err != nil {
//		return "", err
//	}
//	var b []byte
//	//执行AES解密
//	b, err = AesDecrypt(pwdByte, password)
//	if err == nil {
//		return string(b), nil
//	}
//	return "", err
//}
