package function

import (
	"encoding/base64"
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
