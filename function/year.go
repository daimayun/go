package function

// OneYearCycleMonthRange 一年周期月范围
func OneYearCycleMonthRange(start int) (list []int) {
	for i := 1; i <= 12; i++ {
		list = append(list, start)
		if start == 12 {
			start = 1
		} else {
			start++
		}
	}
	return
}
