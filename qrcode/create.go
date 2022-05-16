package qrcode

import (
	"errors"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path"
)

// CreateContainLogoQrCode 生成带有logo图片的二维码
func CreateContainLogoQrCode(content, logoFilePathOrLogoRemoteUrl string, size, logoWidth, logoHeight int) (image.Image, error) {
	if logoWidth == 0 {
		logoWidth = 80
	}
	if logoHeight == 0 {
		logoHeight = 80
	}
	if size == 0 {
		size = 430
	}
	var (
		bgImg        image.Image
		offset       image.Point
		avatarFile   *os.File
		avatarImg    image.Image
		logoFileType string
		err          error
	)
	bgImg, err = createQrCode(content, size)
	if err != nil {
		return nil, errors.New("创建二维码失败[1]: " + err.Error())
	}

	if logoFilePathOrLogoRemoteUrl[0:4] == "http" {
		var response *http.Response
		response, err = http.Get(logoFilePathOrLogoRemoteUrl)
		if err != nil || response.StatusCode != 200 {
			return nil, errors.New("创建二维码失败[2]: 远程图片链接打开失败")
		}
		defer response.Body.Close()
		logoFileType = path.Ext(logoFilePathOrLogoRemoteUrl)
		if logoFileType == ".png" {
			avatarImg, err = png.Decode(response.Body)
		} else if logoFileType == ".jpg" || logoFileType == ".jpeg" {
			avatarImg, err = jpeg.Decode(response.Body)
		} else {
			return nil, errors.New("创建二维码失败[3]: logo图片的后缀不是png|jpg|jpeg")
		}
	} else {
		avatarFile, err = os.Open(logoFilePathOrLogoRemoteUrl)
		if err != nil {
			return nil, errors.New("创建二维码失败[2]: " + err.Error())
		}
		logoFileNameAndSuffix := path.Base(logoFilePathOrLogoRemoteUrl)
		logoFileType = path.Ext(logoFileNameAndSuffix)
		if logoFileType == ".png" {
			avatarImg, err = png.Decode(avatarFile)
		} else if logoFileType == ".jpg" || logoFileType == ".jpeg" {
			avatarImg, err = jpeg.Decode(avatarFile)
		} else {
			return nil, errors.New("创建二维码失败[3]: logo图片的后缀不是png|jpg|jpeg")
		}
	}

	if err != nil {
		return nil, errors.New("创建二维码失败[4]: " + err.Error())
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

// createContainLogoQrCode 生成带有logo图片的二维码
func createContainLogoQrCode(content, logoFilePath string, size, logoWidth, logoHeight int) (image.Image, error) {
	if logoWidth == 0 {
		logoWidth = 80
	}
	if logoHeight == 0 {
		logoHeight = 80
	}
	if size == 0 {
		size = 430
	}
	var (
		bgImg      image.Image
		offset     image.Point
		avatarFile *os.File
		avatarImg  image.Image
		err        error
	)
	bgImg, err = createQrCode(content, size)
	if err != nil {
		return nil, errors.New("创建二维码失败[1]: " + err.Error())
	}
	avatarFile, err = os.Open(logoFilePath)
	if err != nil {
		return nil, errors.New("创建二维码失败[2]: " + err.Error())
	}
	logoFileNameAndSuffix := path.Base(logoFilePath)
	logoFileType := path.Ext(logoFileNameAndSuffix)
	if logoFileType == ".png" {
		avatarImg, err = png.Decode(avatarFile)
	} else if logoFileType == ".jpg" || logoFileType == ".jpeg" {
		avatarImg, err = jpeg.Decode(avatarFile)
	} else {
		return nil, errors.New("创建二维码失败[3]: logo图片的后缀不是png|jpg|jpeg")
	}
	if err != nil {
		return nil, errors.New("创建二维码失败[4]: " + err.Error())
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

// createQrCode 创建二维码
func createQrCode(content string, size int) (img image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, qrcode.Highest)
	if err != nil {
		return nil, errors.New("创建二维码失败[5]: " + err.Error())
	}
	qrCode.DisableBorder = true
	img = qrCode.Image(size)
	return img, nil
}

func imageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
