package http

import (
	"bytes"
	"io/ioutil"
	nh "net/http"
	"net/url"
)

// PostJson post_json
func PostJson(url string, jsonByteBuffer *bytes.Buffer, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/json"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/json"}}
	}
	return Request(nh.MethodPost, url, jsonByteBuffer, headers...)
}

// PostForm post_form
func PostForm(url string, data url.Values) (b []byte, err error) {
	var res *nh.Response
	res, err = nh.PostForm(url, data)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
