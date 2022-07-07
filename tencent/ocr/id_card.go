package ocr

import (
	"context"
	"errors"
	v20181119 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// IdCard 身份证识别
func (c *client) IdCard(request *v20181119.IDCardOCRRequest) (response *v20181119.IDCardOCRResponse, err error) {
	if request == nil {
		request = v20181119.NewIDCardOCRRequest()
	}
	if c.GetCredential() == nil {
		err = errors.New("IDCardOCR require credential")
		return
	}
	request.SetContext(context.Background())
	response = v20181119.NewIDCardOCRResponse()
	err = c.Send(request, response)
	return
}
