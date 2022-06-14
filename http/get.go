package http

import (
	"io/ioutil"
	httpRequest "net/http"
)

// Get get请求
func Get(url string) (b []byte, err error) {
	var res *httpRequest.Response
	res, err = httpRequest.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
