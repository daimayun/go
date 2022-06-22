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

// SliceUnique 切片去重
func SliceUnique(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}
	var (
		newSlice = make([]string, 0)
		m        = make(map[string]interface{})
	)
	for _, v := range slice {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = 0
		newSlice = append(newSlice, v)
	}
	return newSlice
}
