package function

import "regexp"

func CheckMobileNumRule(mobile string) bool {
	re := regexp.MustCompile(`^1[3456789]\d{9}$`)
	return re.MatchString(mobile)
}

func CheckEmailRule(email string) bool {
	re := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	return re.MatchString(email)
}

func CheckIdNumRule(idNum string) bool {
	re := regexp.MustCompile(`^((\d{18})|([0-9x]{18})|([0-9X]{18}))$`)
	return re.MatchString(idNum)
}
