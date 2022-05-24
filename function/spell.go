package function

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// SpellDigitalFormat 拼读格式
type SpellDigitalFormat map[string]string

// replaceStringSpellDigitalFormat 替换字符串拼读格式
var replaceStringSpellDigitalFormat SpellDigitalFormat = map[string]string{
	"0": "0",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"s": "s",
	"b": "b",
	"q": "q",
	"w": "w",
	"y": "y",
	".": "d",
}

// defaultSpellDigitalFormat 系统默认数字大写拼读格式
var defaultSpellDigitalFormat SpellDigitalFormat = map[string]string{
	"0": "零",
	"1": "一",
	"2": "二",
	"3": "三",
	"4": "四",
	"5": "五",
	"6": "六",
	"7": "七",
	"8": "八",
	"9": "九",
	"s": "十",
	"b": "百",
	"q": "千",
	"w": "万",
	"y": "亿",
	".": "点",
}

//读小数或者整数
func spellReadFloatAndInt(s string, format SpellDigitalFormat) (str string, err error) {
	arr := strings.Split(s, ".")
	la := len(arr)
	if la == 2 || la == 1 {
		if la == 2 {
			pStr := arr[1]
			lPStr := len(pStr)
			var zs = format["."]
			for j := 0; j < lPStr; j++ {
				if string(pStr[j]) == "0" {
					zs += format[string(pStr[j])]
				} else {
					str += zs + format[string(pStr[j])]
					zs = ""
				}
			}
			s = arr[0]
		}
	} else {
		err = errors.New("小数点多了")
		return
	}
	var zs string
	zs, err = spellReadInt(s, format)
	if err != nil {
		return
	}
	str = zs + str
	return
}

// 读整数
func spellReadInt(s string, format SpellDigitalFormat) (str string, err error) {
	l := len(s)
	if l > 1 && string(s[0]) == "0" {
		err = errors.New("最高位不能为0")
		return
	} else if l == 1 {
		str = format[s]
		return
	}
	xw := ""
	xy := false
	w := false
	for i := 0; i < l; i++ {
		thenS := string(s[l-1-i])
		if i == 0 {
			//个
			if thenS != "0" {
				str = format[thenS] + str
				xy = true
			}
		} else if i == 1 {
			//十
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["s"] + str
			} else {
				xw = format["0"]
			}
		} else if i == 2 {
			//百
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["b"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 3 {
			//千
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["q"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 4 {
			//万
			str = format["w"] + xw + str
			xw = ""
			if thenS != "0" {
				xy = true
				str = format[thenS] + str
				w = true
			} else {
				xy = false
			}
		} else if i == 5 {
			//十万
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["s"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 6 {
			//百万
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["b"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 7 {
			//千万
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["q"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 8 {
			//亿
			str = format["y"] + xw + str
			xw = ""
			if thenS != "0" {
				xy = true
				str = format[thenS] + str
			} else {
				xy = false
			}
		} else if i == 9 {
			//十亿
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["s"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 10 {
			//百亿
			if thenS != "0" {
				xy = true
				str = format[thenS] + format["b"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = format["0"]
				}
			}
		} else if i == 11 {
			//千亿
			if thenS != "0" {
				str = format[thenS] + format["q"] + str
			}
		} else {
			err = errors.New("位数超过限制")
			return
		}
	}
	if w == false {
		str = strings.Replace(str, format["w"], "", 1)
	}
	if len(str) > 1 {
		strRune := []rune(str)
		if string(strRune[utf8.RuneCountInString(str)-1]) == format["0"] {
			str = string(strRune[:utf8.RuneCountInString(str)-1])
		}
		strRune = []rune(str)
		if string(strRune[0])+string(strRune[1]) == format["1"]+format["s"] {
			str = string(strRune[1:])
		}
	}
	return
}

// SpellReadNum 拼读数
func SpellReadNum(str string, formats ...SpellDigitalFormat) (slice []string, err error) {
	format := defaultSpellDigitalFormat
	if len(formats) > 0 {
		format = MapMerge(defaultSpellDigitalFormat, formats[0], true)
	}
	var s string
	s, err = spellReadFloatAndInt(str, format)
	if err != nil {
		return
	}
	strRune := []rune(s)
	for _, v := range strRune {
		slice = append(slice, string(v))
	}
	return
}
