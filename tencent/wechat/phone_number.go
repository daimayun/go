package wechat

import (
	"encoding/json"
	"github.com/daimayun/go/http"
)

// GetPhoneNumber 获取小程序的用户手机号
func GetPhoneNumber(code, accessToken string) (data ResponsePluginPhoneNumberData, err error) {
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=" + accessToken
	var b, reqParam []byte
	reqParam, err = json.Marshal(struct {
		Code string `json:"code"`
	}{Code: code})
	if err != nil {
		return
	}
	b, err = http.PostJson(url, reqParam)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return
}
