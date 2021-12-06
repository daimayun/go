package function

import "github.com/shopspring/decimal"

// Round 四舍五入 [num为保留的小数点位数]
func Round(val float64, num int32) (res float64, exact bool) {
	res, exact = decimal.NewFromFloat(val).Round(num).Float64()
	return
}
