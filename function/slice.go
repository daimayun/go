package function

// InSlice 判断切片中是否存在某值
func InSlice(slice []string, target string) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}
