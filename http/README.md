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
	"fmt"
    "github.com/daimayun/go/http"
    "strings"
)

func main() {
	type requestBody struct {
		Code string `json:"code"`
	}
	reqParam, err := json.Marshal(struct {
        Code string `json:"code"`
	}{Code: "xxx"})
	if err != nil {
		return
	}
	url := "https://www.baidu.com"
	b, err := http.PostJson(url, strings.NewReader(string(reqParam)))
	fmt.Print(string(b))
}
```

#### 2、POST FORM

```go
package main

import (
	"fmt"
	"github.com/daimayun/go/http"
	"net/url"
	"strings"
)

func main() {
	b, err := http.PostForm("https://www.baidu.com", url.Values{"captcha_id": []string{"DkgMECFm9mzXbdFCZnKx"}, "captcha_value": []string{"5027"}})
	if err != nil {
		return
	}
	fmt.Println(string(b))
}
```
