package ocr

import "fmt"

// IdCard 身份证识别
func IdCard(secretId, secretKey, imageLink string, cardSides ...string) (err error) {
	cardSide := "FRONT"
	if len(cardSides) > 0 {
		cardSide = cardSides[0]
	}
	var param string
	param, err = publicParam(secretId, secretKey)
	if err != nil {
		return
	}
	url := "https://ocr.tencentcloudapi.com/?Action=IDCardOCR&ImageUrl=" + imageLink + "&CardSide=" + cardSide + "&" + param
	fmt.Println(url)
	return
}
