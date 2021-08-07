package qrcode

import (
	"errors"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/png"
	"os"
)

var err error

// CreateLogoQrCode 生成带logo的二维码
func CreateLogoQrCode(content, logoFilePath string, size, logoWidth, logoHeight int) (image.Image, error) {
	if logoWidth == 0 {
		logoWidth = 40
	}
	if logoHeight == 0 {
		logoHeight = 40
	}
	if size == 0 {
		size = 430
	}
	var (
		bgImg      image.Image
		offset     image.Point
		avatarFile *os.File
		avatarImg  image.Image
	)
	bgImg, err = createQrCode(content, size)
	if err != nil {
		return nil, errors.New("创建二维码失败[1]: " + err.Error())
	}
	avatarFile, err = os.Open(logoFilePath)
	if err != nil {
		return nil, errors.New("创建二维码失败[2]: " + err.Error())
	}
	avatarImg, err = png.Decode(avatarFile)
	if err != nil {
		return nil, errors.New("创建二维码失败[3]: " + err.Error())
	}
	avatarImg = imageResize(avatarImg, logoWidth, logoHeight)
	b := bgImg.Bounds()
	// 设置为居中
	offset = image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)
	return m, err
}

// createQrCode 创建二维码
func createQrCode(content string, size int) (img image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, qrcode.Highest)
	if err != nil {
		return nil, errors.New("创建二维码失败[4]: " + err.Error())
	}
	qrCode.DisableBorder = true
	img = qrCode.Image(size)
	return img, nil
}

func imageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
