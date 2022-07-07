package ocr

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type client struct {
	common.Client
}

// NewClient 初始化Client
func NewClient(secretId, secretKey, region string) *client {
	c := &client{}
	c.Init(region).WithSecretId(secretId, secretKey).WithProfile(profile.NewClientProfile())
	return c
}
