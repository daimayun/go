package wechat

// ResponseErrorData 请求返回错误时的数据
type ResponseErrorData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ResponseAccessTokenData 返回access_token数据
type ResponseAccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// ResponsePluginOpenIdData 返回小程序OpenID
type ResponsePluginOpenIdData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	OpenId  string `json:"openpid"`
}

// ResponsePaidUnionIdData 返回小程序UnionID
type ResponsePaidUnionIdData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UnionId string `json:"unionid"`
}

// RequestPluginOpenIdData 获取小程序OpenID
type RequestPluginOpenIdData struct {
	Code string `json:"code"`
}
