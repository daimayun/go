package sms

import "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

func Client(regionId, accessKeyId, accessKeySecret string) (client *dysmsapi.Client, err error) {
	client, err = dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	return
}

func SendSms() {
	//
}
