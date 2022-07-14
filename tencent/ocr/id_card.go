package ocr

import (
	"context"
	"errors"
	"github.com/daimayun/go/function"
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
func (c *client) IdCard(request IdCardRequestData) (response IdCardResponseData, err error) {
	if request.ImageUrl == "" && request.ImageBase64 == "" {
		err = errors.New("missing ID card information")
		return
	}
	if c.GetCredential() == nil {
		err = errors.New("IDCardOCR require credential")
		return
	}
	var (
		req *v20181119.IDCardOCRRequest
		res *v20181119.IDCardOCRResponse
	)
	req = v20181119.NewIDCardOCRRequest()
	if request.ImageUrl != "" {
		req.ImageUrl = function.GetStringPointer(request.ImageUrl)
	}
	if request.ImageBase64 != "" {
		req.ImageBase64 = function.GetStringPointer(request.ImageBase64)
	}
	if request.CardSide != "" {
		req.CardSide = function.GetStringPointer(request.CardSide)
	}
	req.SetContext(context.Background())
	res = v20181119.NewIDCardOCRResponse()
	err = c.Send(req, res)
	if err == nil {
		if *res.Response.IdNum != "" && *res.Response.Name != "" && *res.Response.Sex != "" && *res.Response.Nation != "" && *res.Response.Birth != "" && *res.Response.Address != "" {
			response = IdCardResponseData{
				Name:    *res.Response.Name,
				Sex:     *res.Response.Sex,
				Nation:  *res.Response.Nation,
				Birth:   *res.Response.Birth,
				Address: *res.Response.Address,
				IdNum:   *res.Response.IdNum,
			}
		} else if *res.Response.Authority != "" && *res.Response.ValidDate != "" {
			response = IdCardResponseData{
				Authority: *res.Response.Authority,
				ValidDate: *res.Response.ValidDate,
			}
		} else {
			err = errors.New("ID card ocr fail")
		}
		return
	}
	return
}
