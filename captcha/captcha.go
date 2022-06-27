package captcha

import (
	"github.com/golang/freetype/truetype"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var (
	captchaStore = base64Captcha.DefaultMemStore

	height          int    = 50
	width           int    = 200
	noiseCount      int    = 0
	showLineOptions int    = 0
	language        string = "zh"
	length          int    = 6

	font = []string{"wqy-microhei.ttc"}

	bgColor = &color.RGBA{
		R: 3,
		G: 102,
		B: 214,
		A: 125,
	}
)

type NewCaptchaDrivers struct {
	Type            string
	Height          int
	Width           int
	NoiseCount      int
	ShowLineOptions int
	Language        string
	Length          int
	Source          string
	BgColor         *color.RGBA
	fontsStorage    base64Captcha.FontsStorage
	Fonts           []string
	fontsArray      []*truetype.Font
}

// GetCaptcha 获取图形验证码
func GetCaptcha(driverData NewCaptchaDrivers) (id, b64s string, err error) {
	var driver base64Captcha.Driver
	switch driverData.Type {
	case "audio":
		driver = captchaDriverAudio(driverData)
	case "string":
		driver = captchaDriverString(driverData)
	case "math":
		driver = captchaDriverMath(driverData)
	case "chinese":
		driver = captchaDriverChinese(driverData)
	default:
		driver = captchaDriverDigit(driverData)
	}
	c := base64Captcha.NewCaptcha(driver, captchaStore)
	id, b64s, err = c.Generate()
	return
}

// VerifyCaptcha 校验图形验证码[clears[true:验证通过后删除,false:永不删除]]
func VerifyCaptcha(id, value string, clears ...bool) (res bool) {
	clear := true
	if clears != nil {
		clear = clears[0]
	}
	res = captchaStore.Verify(id, value, false) // 只验证不删除
	if res == true && clear == true {
		go captchaStore.Verify(id, value, true) // 验证通过后且需要清理的则删除
	}
	return
}

func captchaDriverAudio(data NewCaptchaDrivers) *base64Captcha.DriverAudio {
	var (
		captchaLength   = length
		captchaLanguage = language
	)
	if data.Length != 0 {
		captchaLength = data.Length
	}
	if data.Language != "" {
		captchaLanguage = data.Language
	}
	return &base64Captcha.DriverAudio{
		Length:   captchaLength,
		Language: captchaLanguage,
	}
}

func captchaDriverString(data NewCaptchaDrivers) *base64Captcha.DriverString {
	var (
		captchaSource          = "1234567890abcdefghijklmnopqrstuvwxyz"
		captchaHeight          = height
		captchaWidth           = width
		captchaLength          = length
		captchaNoiseCount      = noiseCount
		captchaBgColor         = bgColor
		captchaFont            = font
		captchaShowLineOptions = showLineOptions
	)
	if data.Height != 0 {
		captchaHeight = data.Height
	}
	if data.Width != 0 {
		captchaWidth = data.Width
	}
	if data.NoiseCount != 0 {
		captchaNoiseCount = noiseCount
	}
	if data.BgColor != nil {
		captchaBgColor = data.BgColor
	}
	if data.Fonts != nil {
		captchaFont = data.Fonts
	}
	if data.ShowLineOptions != 0 {
		captchaShowLineOptions = data.ShowLineOptions
	}
	if data.Source != "" {
		captchaSource = data.Source
	}
	if data.Length != 0 {
		captchaLength = data.Length
	}
	return &base64Captcha.DriverString{
		Height:          captchaHeight,
		Width:           captchaWidth,
		NoiseCount:      captchaNoiseCount,
		ShowLineOptions: captchaShowLineOptions,
		Length:          captchaLength,
		Source:          captchaSource,
		BgColor:         captchaBgColor,
		Fonts:           captchaFont,
	}
}

func captchaDriverChinese(data NewCaptchaDrivers) *base64Captcha.DriverChinese {
	var (
		captchaSource          = "一二三四五六七八九十百千万"
		captchaHeight          = height
		captchaWidth           = width
		captchaLength          = length
		captchaNoiseCount      = noiseCount
		captchaBgColor         = bgColor
		captchaFont            = font
		captchaShowLineOptions = showLineOptions
	)
	if data.Height != 0 {
		captchaHeight = data.Height
	}
	if data.Width != 0 {
		captchaWidth = data.Width
	}
	if data.NoiseCount != 0 {
		captchaNoiseCount = noiseCount
	}
	if data.BgColor != nil {
		captchaBgColor = data.BgColor
	}
	if data.Fonts != nil {
		captchaFont = data.Fonts
	}
	if data.ShowLineOptions != 0 {
		captchaShowLineOptions = data.ShowLineOptions
	}
	if data.Source != "" {
		captchaSource = data.Source
	}
	if data.Length != 0 {
		captchaLength = data.Length
	}
	return &base64Captcha.DriverChinese{
		Height:          captchaHeight,
		Width:           captchaWidth,
		Length:          captchaLength,
		NoiseCount:      captchaNoiseCount,
		ShowLineOptions: captchaShowLineOptions,
		Source:          captchaSource,
		BgColor:         captchaBgColor,
		Fonts:           captchaFont,
	}
}

func captchaDriverMath(data NewCaptchaDrivers) *base64Captcha.DriverMath {
	var (
		captchaHeight          = height
		captchaWidth           = width
		captchaNoiseCount      = noiseCount
		captchaBgColor         = bgColor
		captchaFont            = font
		captchaShowLineOptions = showLineOptions
	)
	if data.Height != 0 {
		captchaHeight = data.Height
	}
	if data.Width != 0 {
		captchaWidth = data.Width
	}
	if data.NoiseCount != 0 {
		captchaNoiseCount = noiseCount
	}
	if data.BgColor != nil {
		captchaBgColor = data.BgColor
	}
	if data.Fonts != nil {
		captchaFont = data.Fonts
	}
	if data.ShowLineOptions != 0 {
		captchaShowLineOptions = data.ShowLineOptions
	}
	return &base64Captcha.DriverMath{
		Height:          captchaHeight,
		Width:           captchaWidth,
		NoiseCount:      captchaNoiseCount,
		ShowLineOptions: captchaShowLineOptions,
		BgColor:         captchaBgColor,
		Fonts:           captchaFont,
	}
}

func captchaDriverDigit(data NewCaptchaDrivers) *base64Captcha.DriverDigit {
	var (
		captchaHeight = height
		captchaWidth  = width
		captchaLength = length
	)
	if data.Height != 0 {
		captchaHeight = data.Height
	}
	if data.Width != 0 {
		captchaWidth = data.Width
	}
	if data.Length != 0 {
		captchaLength = data.Length
	}
	return &base64Captcha.DriverDigit{
		Height: captchaHeight,
		Width:  captchaWidth,
		Length: captchaLength,
	}
}
