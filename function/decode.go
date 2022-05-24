package function

import (
	"bytes"
	"encoding/base64"
	"math/big"
	"net/url"
)

// UrlDecode url_decode
func UrlDecode(str string) string {
	data, _ := url.QueryUnescape(str)
	return data
}

// Base64Decode base64_decode
func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

// Base58Decode base58_decode
func Base58Decode(str string) string {
	strByte := []byte(str)
	ret := big.NewInt(0)
	for _, byteElem := range strByte {
		index := bytes.IndexByte(base58, byteElem)
		ret.Mul(ret, big.NewInt(58))
		ret.Add(ret, big.NewInt(int64(index)))
	}
	return string(ret.Bytes())
}
