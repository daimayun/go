package https

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Request 执行请求
func Request(method, url string, body io.Reader, headers ...map[string]string) (b []byte, code int, err error) {
	var (
		req *http.Request
		res *http.Response
	)
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Set(key, value)
		}
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	code = res.StatusCode
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	return
}
