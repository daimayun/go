package function

import (
	"errors"
	"strconv"
	"time"
)

var (
	TimeLayoutYMDHIS string = "20060102150405"
	TimeLayout       string = "2006-01-02 15:04:05"
	TimeLayoutYMD    string = "2006-01-02"
	TimeLayoutYM     string = "2006-01"
)

// TimeVariable 时间变量[有可能是正负数]
var TimeVariable int64

// NewTime 创建一个全局指定的 Time
func NewTime(timeStr string, layouts ...string) (err error) {
	var t time.Time
	t, err = StringToTime(timeStr, layouts...)
	if err != nil {
		return
	}
	TimeVariable = 0
	nt := time.Now()
	if t.Unix() != nt.Unix() {
		TimeVariable = t.Unix() - nt.Unix()
	}
	return
}

// CleanTime 清理指定的时间并校准为服务器的时间
func CleanTime() {
	TimeVariable = 0
}

// TimeNow 当前Time
func TimeNow() time.Time {
	now := time.Now()
	if TimeVariable == 0 {
		return now
	}
	return TimestampToTime(now.Unix() + TimeVariable)
}

// TimeNowUnix 当前时间戳
func TimeNowUnix() int64 {
	return TimeNow().Unix()
}

// TodayStartAndEndTime 今天00:00:00时间和今天23:59:59时间
func TodayStartAndEndTime() (startTime, endTime time.Time) {
	return DayStartAndEndTime()
}

// NowWeekStartAndEndTime 本周一00:00:00时间和本周日23:59:59时间
func NowWeekStartAndEndTime() (startTime, endTime time.Time) {
	return WeekStartAndEndTime()
}

// NowMonthStartAndEndTime 本月1号00:00:00时间和本月末23:59:59时间
func NowMonthStartAndEndTime() (startTime, endTime time.Time) {
	return MonthStartAndEndTime()
}

// NowQuarterStartAndEndTime 本季度1号00:00:00时间和本季度末23:59:59时间
func NowQuarterStartAndEndTime() (startTime, endTime time.Time) {
	return QuarterStartAndEndTime()
}

// NowYearStartAndEndTime 本年1月1号00:00:00时间和本年12月31号23:59:59时间
func NowYearStartAndEndTime() (startTime, endTime time.Time) {
	return YearStartAndEndTime()
}

// DayStartAndEndTime 该天00:00:00时间和该天23:59:59时间
func DayStartAndEndTime(ts ...time.Time) (startTime, endTime time.Time) {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	startTime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	endTime = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
	return
}

// WeekStartAndEndTime 该周一00:00:00时间和该周日23:59:59时间
func WeekStartAndEndTime(ts ...time.Time) (startTime, endTime time.Time) {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	offset := int(time.Monday - t.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}
	lastOffset := int(time.Saturday - t.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastOffset == 6 {
		lastOffset = -1
	}
	startTime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	endTime = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, lastOffset+1)
	return
}

// MonthStartAndEndTime 该月1号00:00:00时间和该月末23:59:59时间
func MonthStartAndEndTime(ts ...time.Time) (startTime, endTime time.Time) {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	startTime = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	endTime = time.Date(t.Year(), t.Month(), startTime.AddDate(0, 1, -1).Day(), 23, 59, 59, 0, time.Local)
	return
}

// QuarterStartAndEndTime 该季度1号00:00:00时间和该季度末23:59:59时间
func QuarterStartAndEndTime(ts ...time.Time) (startTime, endTime time.Time) {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	month := int(t.Month())
	if month >= 1 && month <= 3 {
		startTime = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(t.Year(), 3, 31, 23, 59, 59, 0, time.Local)
	} else if month >= 4 && month <= 6 {
		startTime = time.Date(t.Year(), 4, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(t.Year(), 6, 30, 23, 59, 59, 0, time.Local)
	} else if month >= 7 && month <= 9 {
		startTime = time.Date(t.Year(), 7, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(t.Year(), 9, 30, 23, 59, 59, 0, time.Local)
	} else {
		startTime = time.Date(t.Year(), 10, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(t.Year(), 12, 31, 23, 59, 59, 0, time.Local)
	}
	return
}

// YearStartAndEndTime 该年1月1号00:00:00时间和该年12月31号23:59:59时间
func YearStartAndEndTime(ts ...time.Time) (startTime, endTime time.Time) {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	startTime = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	endTime = time.Date(t.Year(), 12, 31, 23, 59, 59, 0, time.Local)
	return
}

// BeforeSecondTime N秒前的时间
func BeforeSecondTime(seconds ...int64) time.Time {
	var second int64 = 1
	if len(seconds) > 0 {
		second = seconds[0]
	}
	return time.Unix(TimeNowUnix()-second, 0)
}

// AfterSecondTime N秒后的时间
func AfterSecondTime(seconds ...int64) time.Time {
	var second int64 = 1
	if len(seconds) > 0 {
		second = seconds[0]
	}
	return time.Unix(TimeNowUnix()+second, 0)
}

// BeforeDayTime N天前的当前时间
func BeforeDayTime(days ...int) time.Time {
	day := 1
	if len(days) > 0 {
		day = days[0]
	}
	return TimeNow().AddDate(0, 0, -day)
}

// AfterDayTime N天后的当前时间
func AfterDayTime(days ...int) time.Time {
	day := 1
	if len(days) > 0 {
		day = days[0]
	}
	return TimeNow().AddDate(0, 0, day)
}

// BeforeMonthTime N月前的当前时间
func BeforeMonthTime(months ...int) time.Time {
	month := 1
	if len(months) > 0 {
		month = months[0]
	}
	return TimeNow().AddDate(0, -month, 0)
}

// AfterMonthTime N月后的当前时间
func AfterMonthTime(months ...int) time.Time {
	month := 1
	if len(months) > 0 {
		month = months[0]
	}
	return TimeNow().AddDate(0, month, 0)
}

// BeforeYearTime N年前的当前时间
func BeforeYearTime(years ...int) time.Time {
	year := 1
	if len(years) > 0 {
		year = years[0]
	}
	return TimeNow().AddDate(-year, 0, 0)
}

// AfterYearTime N年后的当前时间
func AfterYearTime(years ...int) time.Time {
	year := 1
	if len(years) > 0 {
		year = years[0]
	}
	return TimeNow().AddDate(year, 0, 0)
}

// StringToTime 将字符串转为时间[2021-08-08 08:08:08]
func StringToTime(str string, layouts ...string) (time.Time, error) {
	layout := TimeLayout
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.ParseInLocation(layout, str, time.Local)
}

// TimestampToTime 时间戳转时间
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// TimeByDay 根据天数/号返回当前月的时间
func TimeByDay(day int8, hour, minute, second int8) (t time.Time, err error) {
	if day <= 0 || day > 31 {
		err = errors.New("日期天数不正确")
		return
	}
	if hour < 0 || hour > 24 {
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
	return time.ParseInLocation(TimeLayout, TimeNow().Format(TimeLayoutYM)+"-"+d+" "+h+":"+i+":"+s, time.Local)
}

// FebruaryIsLeapYear 二月是否是闰年
func FebruaryIsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// MonthDays 该月有多少天
func MonthDays(ts ...time.Time) int8 {
	t := TimeNow()
	if len(ts) > 0 {
		t = ts[0]
	}
	month := int8(t.Month())
	type dayType map[int8]struct{}
	day31 := dayType{
		1:  {},
		3:  {},
		5:  {},
		7:  {},
		8:  {},
		10: {},
		12: {},
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	day30 := dayType{
		4:  {},
		6:  {},
		9:  {},
		11: {},
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	if FebruaryIsLeapYear(t.Year()) {
		return 29
	}
	return 28
}
