# 一、图形验证码的生成

```go
package main

import "github.com/daimayun/go/captcha"

// CreateCaptcha 生成图形验证码[默认为数字验证码]
func CreateCaptcha() (id, b64s string, err error) {
	var captchaDriver captcha.NewCaptchaDrivers
	return captcha.GetCaptcha(captchaDriver)
}

// CheckCaptcha 校验验证码
func CheckCaptcha(id, value string) bool {
	return captcha.VerifyCaptcha(id, value)
}
```