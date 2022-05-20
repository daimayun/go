package function

import "time"

var (
	TimeLayoutYMDHIS string = "20060102150405"
	TimeLayout       string = "2006-01-02 15:04:05"
	TimeLayoutYMD    string = "2006-01-02"
)

// TodayDateTimeStart 今日00:00:00时间
func TodayDateTimeStart() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
}

// TodayDateTimeEnd 今日23:59:59时间
func TodayDateTimeEnd() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
}

// NowMonthDateTimeStart 当前月1号00:00:00时间
func NowMonthDateTimeStart() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
}

// NowMonthDateTimeEnd 当前月最后一天的23:59:59时间
func NowMonthDateTimeEnd() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	nowDate := firstOfMonth.AddDate(0, 1, -1)
	return time.Date(currentYear, currentMonth, nowDate.Day(), 23, 59, 59, 0, currentLocation)
}

// NowYearDateTimeStart 当年1月1号00:00:00时间
func NowYearDateTimeStart() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
}

// NowYearDateTimeEnd 当年12月31号的23:59:59时间
func NowYearDateTimeEnd() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), 12, 31, 23, 59, 59, 0, currentTime.Location())
}

// BeforeSecondTime N秒前的时间
func BeforeSecondTime(second int64) time.Time {
	return time.Unix(time.Now().Unix()-second, 0)
}

// AfterSecondTime N秒后的时间
func AfterSecondTime(second int64) time.Time {
	return time.Unix(time.Now().Unix()+second, 0)
}

// BeforeDayTime N天前的时间
func BeforeDayTime(day int) time.Time {
	return time.Now().AddDate(0, 0, -day)
}

// AfterDayTime N天后的时间
func AfterDayTime(day int) time.Time {
	return time.Now().AddDate(0, 0, day)
}

// BeforeMonthTime N月前的时间
func BeforeMonthTime(month int) time.Time {
	return time.Now().AddDate(0, -month, 0)
}

// AfterMonthTime N月后的时间
func AfterMonthTime(month int) time.Time {
	return time.Now().AddDate(0, month, 0)
}

//StringToTime 将字符串转为时间（例如：2021-08-08 08:08:08）
func StringToTime(str string, layouts ...string) (time.Time, error) {
	layout := TimeLayout
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Parse(layout, str)
}

// TimestampToTime 时间戳转日期
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
