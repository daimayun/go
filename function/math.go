package function

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// InterceptDecimal 不四舍五入截取小数点 [n为保留的小数点数]
func InterceptDecimal(f float64, n int) float64 {
	d := float64(1)
	if n > 0 {
		d = math.Pow10(n)
	}
	return math.Trunc(f*d) / d
}

// FloatRound 四舍五入 [n为保留的小数点位数] [不优先使用]
func FloatRound(f float64, n int) (res float64, err error) {
	res, err = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(n)+"f", f+CreateFraction(f, n)), 64)
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
	return strconv.FormatFloat(f+CreateFraction(f, n), 'f', n, 64)
}

// UpInteger 向上取整
func UpInteger(f float64) float64 {
	return math.Ceil(f)
}

// DownInteger 向下取整
func DownInteger(f float64) float64 {
	return math.Floor(f)
}

// UpIntegerToInt64 向上取整返回int64
func UpIntegerToInt64(f float64) int64 {
	return Float64ToInt64(UpInteger(f))
}

// DownIntegerToInt64 向下取整返回int64
func DownIntegerToInt64(f float64) int64 {
	return Float64ToInt64(DownInteger(f))
}

// CreateFraction 创建分数[n为保留的小数点]
func CreateFraction(f64 float64, n int) float64 {
	var f float64 = 0
	slice := strings.Split(Float64ToString(f64), ".")
	if len(slice) == 2 {
		str := strings.TrimRight(slice[1], "0")
		strLen := len(str)
		if strLen > n {
			f = 1 / math.Pow10(strLen+1)
		}
	}
	return f
}
