package ocr

import (
	"encoding/json"
	"errors"
	"github.com/daimayun/go/function"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AliOCRHost                    = "https://personcard.market.alicloudapi.com" // url
	AliOCRPath                    = "/ai-market/ocr/personid"                   // path
	AliOCRImageTypeByFIleContents = "0"                                         // 图片内容类型
	AliOCRImageTypeByFileLink     = "1"                                         // 图片链接地址类型
	AliOCRIdCardSide              = "FRONT"                                     // 无需指定正反面，填写默认值FRONT即可，系统后台自动识别正反面
	AliOCRIdCardImageIsOk         = "1"                                         // 身份证照片信息读取成功
)

// AliOCRResponsePersonIdSideEntityData 正反面图片及个数信息
type aliOCRResponsePersonIdSideEntityData struct {
	PersonIdFrontSideAmount string `json:"PERSON_ID_FRONT_SIDE_AMOUNT"`
	PersonIdFrontSideStatus string `json:"PERSON_ID_FRONT_SIDE_STATUS"`
	PersonIdBackSideAmount  string `json:"PERSON_ID_BACK_SIDE_AMOUNT"`
	PersonIdBackSideStatus  string `json:"PERSON_ID_BACK_SIDE_STATUS"`
}

// AliOCRResponsePersonIdEntityData 图片文字信息
type AliOCRResponsePersonIdEntityData struct {
	PersonName            string `json:"PERSON_NAME"`   // 姓名
	PersonSex             string `json:"PERSON_SEX"`    // 性别
	PersonNation          string `json:"PERSON_NATION"` // 民族
	PersonId              string `json:"PERSON_ID"`     // 公民身份号码
	PersonBirth           string `json:"PERSON_BIRTH"`  // 出生
	PersonChineseBirth    string `json:"PERSON_CHINESE_BIRTH"`
	PersonSign            string `json:"PERSON_SIGN"`
	PersonAnimal          string `json:"PERSON_ANIMAL"`
	PersonBirthCountryId  string `json:"PERSON_BIRTH_COUNTRY_ID"`
	PersonBirthProvince   string `json:"PERSON_BIRTH_PROVINCE"`
	PersonBirthCity       string `json:"PERSON_BIRTH_CITY"`
	PersonBirthDistrict   string `json:"PERSON_BIRTH_DISTRICT"`
	PersonAddress         string `json:"PERSON_ADDRESS"` // 住址
	PersonLivingCountryId string `json:"PERSON_LIVING_COUNTRY_ID"`
	PersonLivingProvince  string `json:"PERSON_LIVING_PROVINCE"`
	PersonLivingCity      string `json:"PERSON_LIVING_CITY"`
	PersonLivingDistrict  string `json:"PERSON_LIVING_DISTRICT"`
	Organization          string `json:"ORGANIZATION"` // 签发机关
	OrganizationCountryId string `json:"ORGANIZATION_COUNTRY_ID"`
	OrganizationProvince  string `json:"ORGANIZATION_PROVINCE"`
	OrganizationCity      string `json:"ORGANIZATION_CITY"`
	OrganizationDistrict  string `json:"ORGANIZATION_DISTRICT"`
	TimeZone              string `json:"TIME_ZONE"` // 有效期限
}

// AliOCRResponseData 阿里云OCR返回数据
type AliOCRResponseData struct {
	PersonIdStatus     string                               `json:"PERSON_ID_STATUS"`
	PersonIdEntity     AliOCRResponsePersonIdEntityData     `json:"PERSON_ID_ENTITY"`
	PersonIdSideEntity aliOCRResponsePersonIdSideEntityData `json:"PERSON_ID_SIDE_ENTITY"`
}

// GetIdCardImageData 获取身份证图片数据信息
func GetIdCardImageData(appCode, fileBase64EncodeData string, imageType ...string) (data AliOCRResponseData, err error) {
	if appCode == "" {
		err = errors.New("缺少APPCODE")
		return
	}
	imgType := AliOCRImageTypeByFIleContents
	if len(imageType) > 0 {
		imgType = imageType[0]
	}
	var (
		req *http.Request
		res *http.Response
		b   []byte
	)
	req, err = http.NewRequest("POST", AliOCRHost+AliOCRPath, strings.NewReader("AI_IDCARD_IMAGE="+function.UrlEncode(fileBase64EncodeData)+"&AI_IDCARD_IMAGE_TYPE="+imgType+"&AI_IDCARD_SIDE="+AliOCRIdCardSide))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Authorization", "APPCODE "+appCode)
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &data)
	return
}
