package sms

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

const endpoint = "dysmsapi.aliyuncs.com"

// RequestSendSmsData 请求发送短信数据
type RequestSendSmsData struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	RegionId        string `json:"region_id,omitempty"`
	PhoneNumbers    string `json:"phone_numbers"`
	SignName        string `json:"sign_name"`
	TemplateCode    string `json:"template_code"`
	TemplateParam   string `json:"template_param"`
}

// CreateClient 创建Client
func CreateClient(accessKeyId, accessKeySecret string) (result *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String(endpoint),
	}
	result = &dysmsapi20170525.Client{}
	result, err = dysmsapi20170525.NewClient(config)
	return
}

// SendSms 发送短信
func SendSms(data RequestSendSmsData) (err error) {
	var client *dysmsapi20170525.Client
	client, err = CreateClient(data.AccessKeyId, data.AccessKeySecret)
	if err != nil {
		return
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(data.PhoneNumbers),
		SignName:      tea.String(data.SignName),
		TemplateCode:  tea.String(data.TemplateCode),
		TemplateParam: tea.String(data.TemplateParam),
	}
	runtime := &util.RuntimeOptions{}
	catchError := func() (e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		_, e = client.SendSmsWithOptions(sendSmsRequest, runtime)
		return
	}()
	if catchError != nil {
		var e = &tea.SDKError{}
		if t, ok := catchError.(*tea.SDKError); ok {
			e = t
		} else {
			e.Message = tea.String(catchError.Error())
		}
		err = errors.New("短信发送失败: " + err.Error())
	}
	return
}
