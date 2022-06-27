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

// ResponsePluginPaidUnionIdData 返回小程序UnionID
type ResponsePluginPaidUnionIdData struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UnionId string `json:"unionid"`
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

// ResponsePluginPhoneNumberData 返回小程序的用户手机号码
type ResponsePluginPhoneNumberData struct {
	ErrCode   int64         `json:"errcode"`
	ErrMsg    string        `json:"errmsg"`
	PhoneInfo PhoneInfoData `json:"phone_info"`
}
