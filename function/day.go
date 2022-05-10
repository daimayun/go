package function

import "time"

// GetMonthDay 获得当前月的初始和结束日期
func GetMonthDay() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetWeekDay 获得当前周的初始和结束日期
func GetWeekDay() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	lastOffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastOffset == 6 {
		lastOffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastOffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetQuarterDay 获得当前季度的初始和结束日期
func GetQuarterDay() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

// GetYearDay 获取当年的初始和结束日期
func GetYearDay() (string, string) {
	year := time.Now().Format("2006")
	firstOfYear := year + "-01-01 00:00:00"
	lastOfYear := year + "-12-31 23:59:59"
	return firstOfYear, lastOfYear
}

// GetDay 获取当天的初始和结束日期
func GetDay() (string, string) {
	day := time.Now().Format("2006-01-02")
	firstOfDay := day + " 00:00:00"
	lastOfDay := day + " 23:59:59"
	return firstOfDay, lastOfDay
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期[两个日期内所有天数]
func GetBetweenDates(startDate, endDate string, layouts ...string) (d []string) {
	layout := TimeLayoutYMD
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	timeFormatTpl := TimeLayout
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return
	}
	date2Str := date2.Format(layout)
	d = append(d, date.Format(layout))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(layout)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return
}

// GetBetweenMonths 根据开始日期和结束日期计算出时间段内所有月份
func GetBetweenMonths(startDate, endDate string, layouts ...string) (d []string) {
	layout := TimeLayoutYMD
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	timeFormatTpl := TimeLayout
	if len(layouts) > 1 {
		timeFormatTpl = layouts[1]
	}
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return
	}
	date2Str := date2.Format(layout)
	d = append(d, date.Format(layout))
	for {
		date = date.AddDate(0, 1, 0)
		dateStr := date.Format(layout)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return
}
