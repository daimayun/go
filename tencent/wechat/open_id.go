package wechat

import (
	"encoding/json"
	"github.com/daimayun/go/http"
	"strings"
)

// GetPluginOpenId 获取小程序的OpenID
func GetPluginOpenId(code, accessToken string) (data ResponsePluginOpenIdData, err error) {
	url := "https://api.weixin.qq.com/wxa/getpluginopenpid?access_token=" + accessToken
	var b, reqParam []byte
	reqParam, err = json.Marshal(struct {
		Code string `json:"code"`
	}{Code: code})
	if err != nil {
		return
	}
	b, err = http.PostJson(url, strings.NewReader(string(reqParam)))
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return
}
