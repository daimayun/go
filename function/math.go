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

// FloatRound 四舍五入 [n为保留的小数点位数] [不优先使用]
func FloatRound(f float64, n int) (res float64, err error) {
	res, err = strconv.ParseFloat(fmt.Sprintf("%." + strconv.Itoa(n) + "f", f), 64)
	return
}

// RoundFloat 四舍五入 [n为保留的小数点位数] [优先使用]
func RoundFloat(f float64, n int) (res float64, err error) {
	res, err = strconv.ParseFloat(RoundToString(f, n), 64)
	return
}

// InterceptDecimalToString 不四舍五入截取小数点后为字符串格式 [n为保留的小数点数]
func InterceptDecimalToString(f float64, n int) string {
	return Float64ToString(InterceptDecimal(f, n))
}

// RoundToString 四舍五入后为字符串格式 [n为保留的小数点位数]
func RoundToString(f float64, n int) string {
	return strconv.FormatFloat(f, 'f', n, 64)
}

// UpInt 向上取整
func UpInt(f float64) float64 {
	return math.Ceil(f)
}

// DownInt 向下取整
func DownInt(f float64) float64 {
	return math.Floor(f)
}
