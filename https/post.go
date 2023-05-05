package https

import (
	"net/http"
	"strings"
)

// PostJson post_json
func PostJson(url, jsonString string, headers ...map[string]string) (b []byte, code int, err error) {
	if len(headers) > 0 {
		if _, ok := headers[0]["Content-Type"]; !ok {
			headers[0]["Content-Type"] = "application/json"
		}
	} else {
		headers = []map[string]string{{"Content-Type": "application/json"}}
	}
	return Request(http.MethodPost, url, strings.NewReader(jsonString), headers...)
}
