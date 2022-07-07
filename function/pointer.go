package function

// GetStringPointer 返回字符串类型的指针
func GetStringPointer(str string) *string {
	return &str
}

// GetStringPointerValue 返回字符串指针类型的值
func GetStringPointerValue(str *string) string {
	return *str
}
