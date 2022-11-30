package function

import (
	"regexp"
	"strconv"
	"strings"
)

// IsChinese 是否全部为中文
func IsChinese(str string) bool {
	return regexp.MustCompile("^[\u4e00-\u9fa5]+$").MatchString(str)
}

// IsCCBCardNo 是否为建设银行卡号 [2020.09.30新增62153403][2019.11.28新增621673][普通高中学生资助卡:623094]
func IsCCBCardNo(cardNo string) bool {
	if regexp.MustCompile(`^\d$`).MatchString(cardNo) == false {
		return false
	}
	list := []string{"621284", "436742", "589970", "620060", "621081", "621467", "621598", "621621", "621700", "622280",
		"622700", "623211", "623668", "621673", "623094", "421349", "434061", "434062", "524094", "526410", "552245",
		"621080", "621082", "621466", "621488", "621499", "622966", "622988", "622382", "621487", "621083", "621084",
		"620107", "62153403", "436742193", "622280193", "5453242", "5491031", "5544033", "622725", "622728", "436728",
		"453242", "491031", "544033", "622707", "625955", "625956", "625362", "625363", "628316", "628317", "356896",
		"356899", "356895", "436718", "436738", "436745", "436748", "489592", "531693", "532450", "532458", "544887",
		"552801", "557080", "558895", "559051", "622166", "622168", "622708", "625964", "625965", "625966", "628266",
		"628366", "622381", "622675", "622676", "622677", "53242", "53243", "553242", "623251", "623669", "624329",
		"623350", "624412", "623644", "62844800", "624458"}
	for _, v := range list {
		if v == "" || v == " " {
			continue
		}
		if strings.Index(cardNo, v) == 0 {
			if len(cardNo) > len(v) {
				return true
			}
		}
	}
	return false
}

// CheckMobileNumRule 验证手机号码
func CheckMobileNumRule(mobile string) bool {
	return regexp.MustCompile(`^1[3456789]\d{9}$`).MatchString(mobile)
}

// CheckEmailRule 验证电子邮箱
func CheckEmailRule(email string) bool {
	return regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(email)
}

// CheckIdNumRuleSimple 简单校验身份证号码的位数+组成字符
func CheckIdNumRuleSimple(idNum string) bool {
	return regexp.MustCompile(`^((\d{18})|([0-9x]{18})|([0-9X]{18}))$`).MatchString(idNum)
}

// CheckIdNumRule 验证身份证号码的合法性
func CheckIdNumRule(idNum string) bool {
	if CheckIdNumRuleSimple(idNum) == false {
		return false
	}
	var (
		idNumByte   [18]byte
		idNumByte17 [17]byte
	)
	for k, v := range []byte(idNum) {
		idNumByte[k] = v
		if k <= 16 {
			idNumByte17[k] = v
		}
	}
	return verifyId(checkId(idNumByte17), byte2int(idNumByte[17]))
}

// CheckWeChatOpenId 验证微信公众号和小程序的OpenID
func CheckWeChatOpenId(openId string) bool {
	return regexp.MustCompile(`^[0-9A-Za-z_-]{28}$`).MatchString(openId)
}

// CheckWeChatUnionId 验证微信公众号和小程序的UnionID
func CheckWeChatUnionId(unionId string) bool {
	return regexp.MustCompile(`^[0-9A-Za-z_-]{28,30}$`).MatchString(unionId)
}

func byte2int(x byte) byte {
	if x == 88 {
		return 'X'
	}
	return x - 48
}

func checkId(id [17]byte) int {
	array := make([]int, 17)
	for index, value := range id {
		array[index], _ = strconv.Atoi(string(value))
	}
	var wi [17]int = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += array[i] * wi[i]
	}
	return res % 11
}

func verifyId(verify int, idByte byte) bool {
	var temp byte
	var i int
	a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}
	for i = 0; i < 11; i++ {
		if i == verify {
			temp = a18[i]
			break
		}
	}
	if temp == idByte {
		return true
	}
	return false
}
