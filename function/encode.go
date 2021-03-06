package function

import (
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/url"
)

// UrlEncode url_encode
func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// Base64Encode base64_encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// JsonEncode json_encode
func JsonEncode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// 标准BTC base58字符顺序[标准base58字符顺序为:123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ]
var base58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base58Encode base58_encode
func Base58Encode(str string) string {
	//1. 转换成ascii码对应的值
	strByte := []byte(str)
	//2. 转换十进制
	strTen := big.NewInt(0).SetBytes(strByte)
	//3. 取出余数
	var modSlice []byte
	for strTen.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0) //余数
		strTen58 := big.NewInt(58)
		strTen.DivMod(strTen, strTen58, mod)             //取余运算
		modSlice = append(modSlice, base58[mod.Int64()]) //存储余数,并将对应值放入其中
	}
	// 处理0就是1的情况 0使用字节'1'代替
	for _, elem := range strByte {
		if elem != 0 {
			break
		} else if elem == 0 {
			modSlice = append(modSlice, byte('1'))
		}
	}
	return string(reverseByteArr(modSlice))
}

// reverseByteArr 将字节的数组反转
func reverseByteArr(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i] //前后交换
	}
	return bytes
}
