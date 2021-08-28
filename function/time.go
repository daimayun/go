package function

import "time"

var (
	TimeLayout string = "2006-01-02 15:04:05"
	TimeLayoutYMD string = "2006-01-02"
	TimeLayoutHMS string = "15:04:05"
)

func NowDataTime() time.Time {
	return time.Now()
}

func NowDataTimeFmtStr() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(TimeLayout)
}

func NowDataTimeStr() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("20060102150405")
}

// TodayDataTimeStart 今日凌晨时间
func TodayDataTimeStart() (timestamp int64, Time time.Time, timeStr string) {
	currentTime := time.Now()
	Time = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	timestamp = Time.Unix()
	timeStr = Time.Format(TimeLayout)
	return
}

// TodayDataTimeEnd 今日23:59:59时间
func TodayDataTimeEnd() (timestamp int64, Time time.Time, timeStr string) {
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
