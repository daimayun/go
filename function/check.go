package function

import (
	"regexp"
	"strconv"
)

// CheckMobileNumRule 验证手机号码
func CheckMobileNumRule(mobile string) bool {
	re := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return re.MatchString(mobile)
}

// CheckEmailRule 验证电子邮箱
func CheckEmailRule(email string) bool {
	re := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	return re.MatchString(email)
}

// CheckIdNumRuleSimple 简单校验身份证号码的位数+组成字符
func CheckIdNumRuleSimple(idNum string) bool {
	re := regexp.MustCompile(`^((\d{18})|([0-9x]{18})|([0-9X]{18}))$`)
	return re.MatchString(idNum)
}

// CheckIdNumRule 验证身份证号码的合法性
func CheckIdNumRule(idNum string) bool {
	re := regexp.MustCompile(`^((\d{18})|([0-9x]{18})|([0-9X]{18}))$`)
	if re.MatchString(idNum) == false {
		return false
	}
	var (
		idNumByte [18]byte
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

func byte2int(x byte) byte {
	if x == 88 {
		return 'X'
	}
	return x - 48
}

func checkId(id [17]byte) int {
	array := make([]int, 17)
	for index , value := range id {
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
