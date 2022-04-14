package function

import (
	"fmt"
	"strconv"
)

// Int64ToInt int64转int
func Int64ToInt(i64 int64) int {
	strInt64 := strconv.FormatInt(i64, 10)
	i, _ := strconv.Atoi(strInt64)
	return i
}

// Int64ToString int64转string
func Int64ToString(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// Float64ToString float64转string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float32ToString float32转string
func Float32ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 32)
}

// StringToInt64 string转int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// StringToFloat64 string转float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float64ToInt64 float64转int64
func Float64ToInt64(f float64) int64 {
	return int64(f)
}

// Int64ToFloat64 int64转float64
func Int64ToFloat64(i64 int64) float64 {
	return float64(i64)
}

// RemoveInvalid0 去掉无效的0
func RemoveInvalid0(f float64) float64 {
	f64, _ := StringToFloat64(RemoveInvalid0ToString(f))
	return f64
}

// RemoveInvalid0ToString 去掉无效的0并转为字符串格式
func RemoveInvalid0ToString(f float64) string {
	return fmt.Sprintf("%g", f)
}
