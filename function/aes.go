package function

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// 只支持16、24、32位，分别对应AES-128，AES-192，AES-256 加密方法
var password = []byte("774D58AB5192D2556F5C1D39C6E049E5")

// AesEncryptToBase64 加密后为Base64格式的字符串
func AesEncryptToBase64(str string, opts ...string) (string, error) {
	res, err := AesEncrypt(str, opts...)
	if err != nil {
		return "", err
	}
	return Base64Encode(res), nil
}

// AesDecryptByBase64 对Base64格式的字符串AES解密
func AesDecryptByBase64(str string, opts ...string) (string, error) {
	str = Base64Decode(str)
	if str == "" {
		return "", errors.New("解密字符串格式不正确")
	}
	res, err := AesDecrypt(str, opts...)
	if err != nil {
		return "", err
	}
	return res, nil
}

// AesEncrypt AES加密
func AesEncrypt(str string, opts ...string) (string, error) {
	mode := "cbc"
	key := password
	lens := len(opts)
	if lens > 0 {
		key = []byte(opts[0])
		if lens > 1 {
			mode = opts[1]
		}
	}
	var (
		err       error
		encrypted []byte
	)
	if mode == "cbc" {
		encrypted, err = AesEncryptByCBC([]byte(str), key)
	} else if mode == "ecb" {
		encrypted = AesEncryptByECB([]byte(str), key)
	} else if mode == "cfb" {
		encrypted, err = AesEncryptByCFB([]byte(str), key)
	} else {
		encrypted, err = AesEncryptByCBC([]byte(str), key)
	}
	if err != nil {
		return "", err
	}
	return Base64Encode(string(encrypted)), err
}

// AesDecrypt AES解密
func AesDecrypt(str string, opts ...string) (string, error) {
	mode := "cbc"
	key := password
	lens := len(opts)
	if lens > 0 {
		key = []byte(opts[0])
		if lens > 1 {
			mode = opts[1]
		}
	}
	var (
		err       error
		decrypted []byte
	)
	if mode == "cbc" {
		decrypted, err = AesDecryptByCBC([]byte(str), key)
	} else if mode == "ecb" {
		decrypted = AesDecryptByECB([]byte(str), key)
	} else if mode == "cfb" {
		decrypted, err = AesDecryptByCFB([]byte(str), key)
	} else {
		decrypted, err = AesDecryptByCBC([]byte(str), key)
	}
	if err != nil {
		return "", err
	}
	return Base64Encode(string(decrypted)), err
}

// AesEncryptByCBC AES加密[CBC模式]
func AesEncryptByCBC(origData []byte, key []byte) (encrypted []byte, err error) {
	//创建加密算法实例
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted = make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(encrypted, origData)
	return
}

// AesDecryptByCBC AES解密[CBC模式]
func AesDecryptByCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	//创建加密算法实例
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted = make([]byte, len(encrypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(decrypted, encrypted)
	//去除填充字符串
	decrypted, err = PKCS7UnPadding(decrypted)
	return
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

// AesEncryptByCFB AES加密[CFB模式]
func AesEncryptByCFB(origData []byte, key []byte) (encrypted []byte, err error) {
	key, _ = hex.DecodeString(string(key))
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return
}

// AesDecryptByCFB AES解密[CFB模式]
func AesDecryptByCFB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	key, err = hex.DecodeString(string(key))
	if err != nil {
		return
	}
	decrypted, err = hex.DecodeString(string(encrypted))
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(decrypted) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}
	iv := decrypted[:aes.BlockSize]
	decrypted = decrypted[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)
	return
}

// AesEncryptByOFB AES加密[OFB模式]
func AesEncryptByOFB(origData []byte, key []byte) (encrypted []byte, err error) {
	origData = PKCS7Padding(origData, aes.BlockSize)
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return
}

// AesDecryptByOFB AES解密[OFB模式]
func AesDecryptByOFB(data []byte, key []byte) (decrypted []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		err = errors.New("data is not a multiple of the block size")
		return
	}
	decrypted = make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decrypted, data)
	decrypted, err = PKCS7UnPadding(decrypted)
	return
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
