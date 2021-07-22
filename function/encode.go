package function

import (
	Url "net/url"
)

// UrlEncode url_encode
func UrlEncode(url string) string {
	return Url.QueryEscape(url)
}
