package wechat

import (
	"encoding/json"
	"github.com/daimayun/go/http"
)

// GetPaidUnionId 获取小程序的UnionID
func GetPaidUnionId(openId, accessToken string) (data ResponsePluginPaidUnionIdData, err error) {
	url := "https://api.weixin.qq.com/wxa/getpaidunionid?access_token=" + accessToken + "&openid=" + openId
	var b []byte
	b, err = http.Get(url)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return
}
