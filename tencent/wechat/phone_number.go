package wechat

import (
	"encoding/json"
	"github.com/daimayun/go/http"
	"strings"
)

// GetPhoneNumber 获取小程序的用户手机号
func GetPhoneNumber(code, accessToken string) (data ResponsePluginPhoneNumberData, err error) {
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=" + accessToken
	var b, reqParam []byte
	type requestBody struct {
		Code string `json:"code"`
	}
	reqParam, err = json.Marshal(requestBody{Code: code})
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
