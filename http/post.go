package http

import (
	"bytes"
	"io/ioutil"
	nh "net/http"
	"net/url"
)

// postJson post_json
func postJson(url string, jsonByteBuffer *bytes.Buffer, headers ...map[string]string) (b []byte, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/json"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/json"}}
	}
	return Request(nh.MethodPost, url, jsonByteBuffer, headers...)
}

// PostJson post_json
func PostJson(url string, jsonByte []byte) (b []byte, err error) {
	var (
		req *nh.Request
		res *nh.Response
	)
	req, err = nh.NewRequest(nh.MethodPost, url, bytes.NewBuffer(jsonByte))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = nh.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
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
