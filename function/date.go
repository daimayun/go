package function

import (
	"errors"
	"strconv"
	"time"
)

// SubMonths 两个时间相隔多少个月[t1减t2][优先使用]
func SubMonths(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	chaY := y1 - y2
	chaM := 12 - m2 + m1
	month = chaM + ((chaY - 1) * 12)
	return
}

// subMonths 两个时间相隔多少个月[t1减t2]
func subMonths(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()
	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

// SubDays 两个时间相隔多少天[t1减t2]
func SubDays(t1, t2 time.Time) (day int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t_ := t1
		t1 = t2
		t2 = t_
		swap = true
	}
	day = int(t1.Sub(t2).Hours() / 24)
	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		day += 1
	}
	if swap {
		day = -day
	}
	return
}

// GetBetweenDays 根据开始日期和结束日期计算出时间段内所有日期[两个日期内所有天数]
func GetBetweenDays(startDate, endDate string, layouts ...string) (d []string) {
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
	var date2 time.Time
	date2, err = time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	// 日期相等直接返回
	if startDate == endDate {
		d = []string{startDate}
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
	var date2 time.Time
	date2, err = time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return
	}
	// 日期相等直接返回
	if startDate == endDate {
		d = []string{startDate}
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

// DayToTime 根据天数/号返回指定 Time
func DayToTime(day int8, hour, minute, second int8, ts ...time.Time) (t time.Time, err error) {
	ti := TimeNow()
	if len(ts) > 0 {
		ti = ts[0]
	}
	if day <= 0 || day > 31 {
		err = errors.New("日期天数不正确")
		return
	}
	if hour < 0 || hour > 23 {
		err = errors.New("日期小时不正确")
		return
	}
	if minute < 0 || minute > 59 {
		err = errors.New("日期分钟不正确")
		return
	}
	if second < 0 || second > 59 {
		err = errors.New("日期秒钟不正确")
		return
	}
	d := strconv.Itoa(int(day))
	if day < 10 {
		d = "0" + d
	}
	h := strconv.Itoa(int(hour))
	if hour < 10 {
		h = "0" + h
	}
	i := strconv.Itoa(int(minute))
	if minute < 10 {
		i = "0" + i
	}
	s := strconv.Itoa(int(second))
	if second < 10 {
		s = "0" + s
	}
	return time.ParseInLocation(TimeLayout, ti.Format(TimeLayoutYM)+"-"+d+" "+h+":"+i+":"+s, time.Local)
}
