package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetAccessToken 获取AccessToken
func GetAccessToken(appId, appSecret string) (data ResponseReturnAccessTokenData, errData ResponseReturnErrorData, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	var (
		res *http.Response
		b   []byte
	)
	res, err = http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
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
