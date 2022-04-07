package function

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

// Round 四舍五入 [num为保留的小数点位数]
func Round(val float64, num int32) (res float64, exact bool) {
	res, exact = decimal.NewFromFloat(val).Round(num).Float64()
	return
}

// Intercept 截取小数点 [retain为保留的小数点数]
func Intercept(val float64, retain int32) (res float64) {
	return
}

// FloatRound 四舍五入 [n为保留的小数点位数]
func FloatRound(f float64, n int) (res float64, err error) {
	format := "%." + strconv.Itoa(n) + "f"
	res, err = strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return
}
