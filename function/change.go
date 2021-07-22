package function

import (
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
	return strconv.Itoa(Int64ToInt(i64))
}

// Float64ToString float64转string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// Float32ToString float32转string
func Float32ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 32)
}

// StrToInt64 string转int6
func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
