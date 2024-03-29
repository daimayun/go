package function

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

// MoreFloat64Mul 多个Float64相乘
func MoreFloat64Mul(fn ...float64) (res float64, exact bool) {
	fnLength := len(fn)
	if fnLength == 0 {
		return
	}
	if fnLength == 1 {
		return fn[0], true
	}
	v := decimal.NewFromFloat(fn[0])
	for i := 1; i < fnLength; i++ {
		v = v.Mul(decimal.NewFromFloat(fn[i]))
	}
	res, exact = v.Float64()
	return
}

// Float64MulFloat64 float64 * float64
func Float64MulFloat64(f1, f2 float64) (res float64, exact bool) {
	res, exact = decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return
}

// Float64MulInt64 float64 * int64
func Float64MulInt64(f float64, i int64) (res float64, exact bool) {
	res, exact = decimal.NewFromFloat(f).Mul(decimal.NewFromInt(i)).Float64()
	return
}

// InterceptDecimal 不四舍五入截取小数点 [n为保留的小数点数位数]
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

// InterceptDecimalToString 不四舍五入截取小数点后为字符串格式 [n为保留的小数点数位数]
func InterceptDecimalToString(f float64, n int) string {
	return Float64ToString(InterceptDecimal(f, n))
}

// RoundToString 四舍五入后为字符串格式 [n为保留的小数点位数]
func RoundToString(f float64, n int) string {
	return strconv.FormatFloat(f+CreateFraction(f, n), 'f', n, 64)
}

// UpInteger 向上取整
func UpInteger(f float64) float64 {
	return Ceil(f)
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

// CreateFraction 创建分数[n为保留的小数点位数]
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

// Ceil 向下取值/整[n为保留的小数点位数]
func Ceil(f64 float64, n ...int) float64 {
	var nums int
	if len(n) > 0 {
		nums = n[0]
	}
	p := math.Pow10(nums)
	return math.Ceil(f64*p) / p
}
