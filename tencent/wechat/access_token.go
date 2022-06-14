package wechat

import (
	"encoding/json"
	"github.com/daimayun/go/http"
)

// GetAccessToken 获取AccessToken[包含小程序和微信公众号]
func GetAccessToken(appId, appSecret string) (data ResponseAccessTokenData, errData ResponseErrorData, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	var b []byte
	b, err = http.Get(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	if data.AccessToken != "" && data.ExpiresIn > 0 {
		return
	}
	err = json.Unmarshal(b, &errData)
	return
}
