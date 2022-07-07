package ocr

import (
	"context"
	"errors"
	v20181119 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// IdCardRequestData OCR身份证识别请求参数
type IdCardRequestData struct {
	ImageUrl    string `json:"image_url,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
	CardSide    string `json:"card_side,omitempty"`
}

// IdCardResponseData OCR身份证识别返回信息
type IdCardResponseData struct {
	Name      string `json:"name,omitempty"`       // 姓名（人像面）
	Sex       string `json:"sex,omitempty"`        // 性别（人像面）
	Nation    string `json:"nation,omitempty"`     // 民族（人像面）
	Birth     string `json:"birth,omitempty"`      // 出生日期（人像面）
	Address   string `json:"address,omitempty"`    // 地址（人像面）
	IdNum     string `json:"id_num,omitempty"`     // 身份证号（人像面）
	Authority string `json:"authority,omitempty"`  // 发证机关（国徽面）
	ValidDate string `json:"valid_date,omitempty"` // 证件有效期（国徽面）
}

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
