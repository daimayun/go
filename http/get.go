package http

import (
	"io/ioutil"
	nh "net/http"
)

// Get get请求
func Get(url string) (b []byte, err error) {
	var res *nh.Response
	res, err = nh.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
