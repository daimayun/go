package function

import "time"

var (
	TimeLayoutYMDHIS string = "20060102150405"
	TimeLayout       string = "2006-01-02 15:04:05"
	TimeLayoutYMD    string = "2006-01-02"
)

// 当前时间
var now = time.Now()

// TodayStartAndEndTime 今天00:00:00时间和今天23:59:59时间
func TodayStartAndEndTime() (startTime, endTime time.Time) {
	startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return
}

// NowWeekStartAndEndTime 本周一00:00:00时间和本周日23:59:59时间
func NowWeekStartAndEndTime() (startTime, endTime time.Time) {
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
	startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, lastOffset+1)
	return
}

// NowMonthStartAndEndTime 本月1号00:00:00时间和本月末23:59:59时间
func NowMonthStartAndEndTime() (startTime, endTime time.Time) {
	startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endTime = time.Date(now.Year(), now.Month(), startTime.AddDate(0, 1, -1).Day(), 23, 59, 59, 0, now.Location())
	return
}

// NowQuarterStartAndEndTime 本季度1号00:00:00时间和本季度末23:59:59时间
func NowQuarterStartAndEndTime() (startTime, endTime time.Time) {
	month := int(now.Month())
	if month >= 1 && month <= 3 {
		startTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(now.Year(), 3, 31, 23, 59, 59, 0, time.Local)
	} else if month >= 4 && month <= 6 {
		startTime = time.Date(now.Year(), 4, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(now.Year(), 6, 30, 23, 59, 59, 0, time.Local)
	} else if month >= 7 && month <= 9 {
		startTime = time.Date(now.Year(), 7, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(now.Year(), 9, 30, 23, 59, 59, 0, time.Local)
	} else {
		startTime = time.Date(now.Year(), 10, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(now.Year(), 12, 31, 23, 59, 59, 0, time.Local)
	}
	return
}

// NowYearStartAndEndTime 本年1月1号00:00:00时间和本年12月31号23:59:59时间
func NowYearStartAndEndTime() (startTime, endTime time.Time) {
	startTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	endTime = time.Date(now.Year(), 12, 31, 23, 59, 59, 0, now.Location())
	return
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

// BeforeYearTime N年前的时间
func BeforeYearTime(year int) time.Time {
	return time.Now().AddDate(-year, 0, 0)
}

// AfterYearTime N年后的时间
func AfterYearTime(year int) time.Time {
	return time.Now().AddDate(year, 0, 0)
}

//StringToTime 将字符串转为时间[2021-08-08 08:08:08]
func StringToTime(str string, layouts ...string) (time.Time, error) {
	layout := TimeLayout
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Parse(layout, str)
}

// TimestampToTime 时间戳转时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
