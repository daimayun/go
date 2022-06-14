# HTTP

### 一、GET请求

```go
package main

import "github.com/daimayun/go/http"

func main() {
    http.Get("https://www.baidu.com")
}
```

### 二、POST请求

#### 1、POST JSON
```go
package main

import (
    "encoding/json"
    "github.com/daimayun/go/http"
    "strings"
)

func main() {
	type requestBody struct {
		Code string `json:"code"`
	}
	reqParam, err = json.Marshal(requestBody{Code: code})
	if err != nil {
		return
	}
	b, err = http.PostJson(url, strings.NewReader(string(reqParam)))
}
```
