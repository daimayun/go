package function

// MapMerge 两个MAP合并,两个MAP键相同的情况,后面MAP的值覆盖前面MAP的值
func MapMerge(m1, m2 map[string]string) (m map[string]string) {
	m = make(map[string]string)
	for i, v := range m1 {
		for j, w := range m2 {
			if i == j {
				m[i] = w
			} else {
				if _, ok := m[i]; !ok {
					m[i] = v
				}
				if _, ok := m[j]; !ok {
					m[j] = w
				}
			}
		}
	}
	return
}
