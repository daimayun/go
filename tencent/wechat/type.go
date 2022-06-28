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
