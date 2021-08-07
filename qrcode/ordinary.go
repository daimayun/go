package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// CreateQrCode 创建普通二维码
func CreateQrCode(content string, width, height int) (image barcode.Barcode, err error) {
	if width == 0 {
		width = 430
	}
	if height == 0 {
		height = 430
	}
	image, err = qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return
	}
	image, err = barcode.Scale(image, width, height)
	return
}
