package wechat

// ResponseReturnErrorData 请求返回错误时的数据
type ResponseReturnErrorData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ResponseReturnAccessTokenData 返回access_token数据
type ResponseReturnAccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
