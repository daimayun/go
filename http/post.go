package http

import (
	"io"
	"io/ioutil"
	httpRequest "net/http"
	"net/url"
)

// POST method
const POST = "POST"

// PostJson post_json
func PostJson(url string, jsonStrReader io.Reader, headers ...map[string]string) (b []byte, err error) {
	var (
		req *httpRequest.Request
		res *httpRequest.Response
	)
	req, err = httpRequest.NewRequest(POST, url, jsonStrReader)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Add(key, value)
		}
	}
	res, err = httpRequest.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}

// PostForm post_form
func PostForm(url string, data url.Values) (b []byte, err error) {
	var res *httpRequest.Response
	res, err = httpRequest.PostForm(url, data)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
