package function

import "time"

var TimeLayout string = "2006-01-02 15:04:05"

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
	Time = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 58, 0, currentTime.Location())
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
