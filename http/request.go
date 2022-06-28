package http

import (
	"bytes"
	"io/ioutil"
	nh "net/http"
)

func Request(method, url string, jsonByteBuffer *bytes.Buffer, headers ...map[string]string) (b []byte, err error) {
	var (
		req *nh.Request
		res *nh.Response
	)
	req, err = nh.NewRequest(method, url, jsonByteBuffer)
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Set(key, value)
		}
	}
	res, err = nh.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
