package miniprogram

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

type WatermarkData struct {
	Timestamp int64  `json:"timestamp"`
	Appid     string `json:"appid"`
}

type PhoneInfoData struct {
	PhoneNumber     string        `json:"phoneNumber"`
	PurePhoneNumber string        `json:"purePhoneNumber"`
	CountryCode     string        `json:"countryCode"`
	Watermark       WatermarkData `json:"watermark"`
}

// ResponsePhoneNumberData 返回小程序的用户手机号码
type ResponsePhoneNumberData struct {
	ErrCode   int64         `json:"errcode"`
	ErrMsg    string        `json:"errmsg"`
	PhoneInfo PhoneInfoData `json:"phone_info"`
}

// ResponseCode2SessionData 返回小程序登录凭证校验
type ResponseCode2SessionData struct {
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	SessionKey string `json:"session_key"`
}
