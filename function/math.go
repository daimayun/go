package function

import (
	"fmt"
	"math"
	"strconv"
)

// InterceptDecimal 不四舍五入截取小数点 [n为保留的小数点数]
func InterceptDecimal(f float64, n int) float64 {
	d := float64(1)
	if n > 0 {
		d = math.Pow10(n)
	}
	return math.Trunc(f*d)/d
}

// FloatRound 四舍五入 [n为保留的小数点位数]
func FloatRound(f float64, n int) (res float64, err error) {
	format := "%." + strconv.Itoa(n) + "f"
	res, err = strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return
}

// RoundFloat 四舍五入 [n为保留的小数点位数]
func RoundFloat(f float64, n int) (res float64, err error) {
	res, err = strconv.ParseFloat(strconv.FormatFloat(f, 'f', n, 64), 64)
	return
}

// UpInt 向上取整
func UpInt(f float64) float64 {
	return math.Ceil(f)
}

// DownInt 向下取整
func DownInt(f float64) float64 {
	return math.Floor(f)
}
