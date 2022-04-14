package function

import "time"

var (
	TimeLayout    string = "2006-01-02 15:04:05"
	TimeLayoutYMD string = "2006-01-02"
	TimeLayoutHMS string = "15:04:05"
)

// NowDateTime 当前时间戳
func NowDateTime() time.Time {
	return time.Now()
}

// NowDateTimeFmtStr 当前时间格式
func NowDateTimeFmtStr() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(TimeLayout)
}

// NowDateTimeStr 当前时间字符串
func NowDateTimeStr() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("20060102150405")
}

// TodayDateTimeStart 今日00:00:00时间
func TodayDateTimeStart() (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	timestamp = Time.Unix()
	timeStr = Time.Format(TimeLayout)
	return
}

// TodayDateTimeEnd 今日23:59:59时间
func TodayDateTimeEnd() (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	timestamp = Time.Unix()
	timeStr = Time.Format(TimeLayout)
	return
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

// SubMonths 两个时间相隔多少个月 [t1 - t2]
func SubMonths(t1, t2 time.Time) (month int) {
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

// SubDays 两个时间相隔多少天 [t1 - t2]
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
