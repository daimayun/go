package http

import (
	nh "net/http"
)

// Get get请求
func Get(url string, headers ...map[string]string) (b []byte, err error) {
	return Request(nh.MethodGet, url, nil, headers...)
}
