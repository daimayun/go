# 一、二维码的生成

```go
package main

import (
    "github.com/daimayun/go/qrcode"
    "image/png"
    "os"
)

// CreateQrCode 创建普通二维码
func CreateQrCode(fileName, url string, width, height int) (err error) {
	qrCode, err := qrcode.CreateQrCode(url, width, height)
	if err != nil {
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, qrCode)
	return
}

// CreateContainLogoQrCode 创建带有logo的二维码
func CreateContainLogoQrCode(fileName, url, logoFilePath string, size, logoWidth, logoHeight int) (err error) {
	qrCode, err := qrcode.CreateContainLogoQrCode(url, logoFilePath, size, logoWidth, logoHeight)
	if err != nil {
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	err = png.Encode(file, qrCode)
	return
}
```