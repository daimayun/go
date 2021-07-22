package function

import Url "net/url"

// UrlDecode url_decode
func UrlDecode(str string) (url string) {
	url, _ = Url.QueryUnescape(str)
	return
}
