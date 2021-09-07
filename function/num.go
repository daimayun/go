package function

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var strNum map[string]string = map[string]string{
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

//读小数或者整数
func spellReadFloatAndInt(s string) (str string, err error) {
	arr := strings.Split(s, ".")
	la := len(arr)
	if la == 2 || la == 1 {
		if la == 2 {
			pStr := arr[1]
			lPStr := len(pStr)
			var zs = strNum["."]
			for j := 0; j < lPStr; j++ {
				if string(pStr[j]) == "0" {
					zs += strNum[string(pStr[j])]
				} else {
					str += zs + strNum[string(pStr[j])]
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
	zs, err = spellReadInt(s)
	if err != nil {
		return
	}
	str = zs + str
	return
}

// 读整数
func spellReadInt(s string) (str string, err error) {
	l := len(s)
	if l > 1 && string(s[0]) == "0" {
		err = errors.New("最高位不能为0")
		return
	} else if l == 1 {
		str = strNum[s]
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
				str = strNum[thenS] + str
				xy = true
			}
		} else if i == 1 {
			//十
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["s"] + str
			} else {
				xw = strNum["0"]
			}
		} else if i == 2 {
			//百
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["b"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 3 {
			//千
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["q"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 4 {
			//万
			str = strNum["w"] + xw + str
			xw = ""
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + str
				w = true
			} else {
				xy = false
			}
		} else if i == 5 {
			//十万
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["s"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 6 {
			//百万
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["b"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 7 {
			//千万
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["q"] + xw + str
				xw = ""
				w = true
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 8 {
			//亿
			str = strNum["y"] + xw + str
			xw = ""
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + str
			} else {
				xy = false
			}
		} else if i == 9 {
			//十亿
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["s"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 10 {
			//百亿
			if thenS != "0" {
				xy = true
				str = strNum[thenS] + strNum["b"] + xw + str
				xw = ""
			} else {
				if xw == "" && xy == true {
					xw = strNum["0"]
				}
			}
		} else if i == 11 {
			//千亿
			if thenS != "0" {
				str = strNum[thenS] + strNum["q"] + str
			}
		} else {
			err = errors.New("位数超过限制")
			return
		}
	}
	if w == false {
		str = strings.Replace(str, strNum["w"], "", 1)
	}
	if len(str) > 1 {
		strRune := []rune(str)
		if string(strRune[utf8.RuneCountInString(str)-1]) == strNum["0"] {
			str = string(strRune[:utf8.RuneCountInString(str)-1])
		}
		strRune = []rune(str)
		if string(strRune[0])+string(strRune[1]) == strNum["1"]+strNum["s"] {
			str = string(strRune[1:])
		}
	}
	return
}

// SpellReadNum 拼读数
func SpellReadNum(str string) (slice []string, err error) {
	var s string
	s, err = spellReadFloatAndInt(str)
	if err != nil {
		return
	}
	strRune := []rune(s)
	for _, v := range strRune {
		slice = append(slice, string(v))
	}
	return
}
