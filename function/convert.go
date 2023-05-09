package function

// StringYuanToInt64Fen 字符串元转int64分
func StringYuanToInt64Fen(yuan string) (fen int64, err error) {
	var float float64
	float, err = StringToFloat64(yuan)
	if err != nil {
		return
	}
	float, _ = Float64MulInt64(float, 100)
	fen, err = StringToInt64(InterceptDecimalToString(float, 0))
	return
}

// YuanToInt64Fen 元转分
func YuanToInt64Fen(yuan float64) (fen int64, err error) {
	yuan, _ = Float64MulInt64(yuan, 100)
	fen, err = StringToInt64(InterceptDecimalToString(yuan, 0))
	return
}
