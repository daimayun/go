package function

import (
	"encoding/base64"
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
