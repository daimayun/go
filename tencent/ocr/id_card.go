package ocr

import (
	errs "errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// 公共请求参数
func publicParam(secretId, secretKey string) (js string, err error) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(common.NewCredential(secretId, secretKey), "", cpf)
	var response *ocr.IDCardOCRResponse
	response, err = client.IDCardOCR(ocr.NewIDCardOCRRequest())
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		err = errs.New("An API error has returned: " + err.Error())
		return
	}
	if err != nil {
		return
	}
	js = response.ToJsonString()
	return
}
