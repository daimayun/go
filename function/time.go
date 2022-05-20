package function

import "time"

var (
	TimeLayoutYMDHIS string = "20060102150405"
	TimeLayout       string = "2006-01-02 15:04:05"
	TimeLayoutYMD    string = "2006-01-02"
	TimeLayoutHMS    string = "15:04:05"
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

// NowYearDateTimeStart 当前年1月1号00:00:00时间
func NowYearDateTimeStart() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
}

// NowYearDateTimeEnd 当前年12月31号的23:59:59时间
func NowYearDateTimeEnd() time.Time {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), 12, 31, 23, 59, 59, 0, currentTime.Location())
}

//TimeStrToTime 将时间字符串转为时间戳和time.Time（例如：2021-08-08 08:08:08）
func TimeStrToTime(str string) (timestamp int64, Time time.Time) {
	loc, _ := time.LoadLocation("Local")
	Time, _ = time.ParseInLocation(TimeLayout, str, loc)
	timestamp = Time.Unix()
	return
}

// AddSecondToTime 在当前时间戳的基础上加上秒数，重新生成时间日期
func AddSecondToTime(second int64) (timestamp int64, Time time.Time, timeStr string) {
	timestamp = time.Now().Unix() + second
	Time = time.Unix(timestamp, 0)
	timeStr = Time.Format(TimeLayout)
	return
}

// TimestampToTime 时间戳转日期格式
func TimestampToTime(timestamp int64) (Time time.Time, timeStr string) {
	Time = time.Unix(timestamp, 0)
	timeStr = Time.Format(TimeLayout)
	return
}

// NowTime 当前时间
func NowTime() (timestamp int64, Time time.Time, timeStr string) {
	timestamp = time.Now().Unix()
	Time = time.Unix(timestamp, 0)
	timeStr = Time.Format(TimeLayout)
	return
}

// TimeSplit 时间格式拆分
func TimeSplit(t time.Time) (ymd, hms string) {
	ymd = t.Format(TimeLayoutYMD)
	hms = t.Format(TimeLayoutHMS)
	return
}

// NowTimeSplit 当前时间格式拆分
func NowTimeSplit() (ymd, hms string) {
	var Time time.Time
	Time = time.Unix(time.Now().Unix(), 0)
	ymd = Time.Format(TimeLayoutYMD)
	hms = Time.Format(TimeLayoutHMS)
	return
}

// BeforeDayTime N天前的时间
func BeforeDayTime(day int) (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = currentTime.AddDate(0, 0, -day)
	timeStr = Time.Format(TimeLayout)
	timestamp = Time.Unix()
	return
}

// AfterDayTime N天后的时间
func AfterDayTime(day int) (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = currentTime.AddDate(0, 0, day)
	timeStr = Time.Format(TimeLayout)
	timestamp = Time.Unix()
	return
}

// BeforeMonthTime N月前的时间
func BeforeMonthTime(month int) (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = currentTime.AddDate(0, -month, 0)
	timeStr = Time.Format(TimeLayout)
	timestamp = Time.Unix()
	return
}

// AfterMonthTime N月后的时间
func AfterMonthTime(month int) (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = currentTime.AddDate(0, month, 0)
	timeStr = Time.Format(TimeLayout)
	timestamp = Time.Unix()
	return
}

// DiyTimeFmtStr Diy时间格式
func DiyTimeFmtStr(format string, timestamps ...int64) string {
	var timestamp int64
	if len(timestamps) == 0 {
		timestamp = time.Now().Unix()
	} else {
		timestamp = timestamps[0]
	}
	return time.Unix(timestamp, 0).Format(format)
}
