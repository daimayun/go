package miniprogram

import (
	"encoding/json"
	"github.com/daimayun/go/http"
)

// GetCode2Session 登录凭证校验
func GetCode2Session(appId, appSecret, jsCode string, grantTypes ...string) (data ResponseCode2SessionData, err error) {
	grantType := "authorization_code"
	if len(grantTypes) > 0 {
		grantType = grantTypes[0]
	}
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + jsCode + "&grant_type=" + grantType
	var b []byte
	b, err = http.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return
}
