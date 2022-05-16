package function

import (
	"errors"
	"strings"
	"time"
)

// IdCardBirthHandle 身份证出生日期处理[2020年1月1日]
func IdCardBirthHandle(birth string, layouts ...string) (time.Time, error) {
	layout := "2006年1月2日"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Parse(layout, birth)
}

// IdCardHandle 身份证有效期处理[2018.04.23-2028.04.23]
func IdCardHandle(date string, layouts ...string) (starTime, endTime time.Time, err error) {
	layout := "2006.01.02"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	arr := strings.Split(date, "-")
	if len(arr) != 2 {
		err = errors.New("身份证有效期限格式不正确-1")
		return
	}
	starTime, err = time.Parse(layout, arr[0])
	if err != nil {
		return
	}
	endTime, err = time.Parse(layout, arr[1])
	if err != nil {
		return
	}
	if endTime.Unix() <= starTime.Unix() {
		err = errors.New("身份证有效期限格式不正确-2")
		return
	}
	return
}
