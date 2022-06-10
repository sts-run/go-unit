package go_unit

import "time"

var (
	templateOne   = "2006-01-02T15:04:05+08:00"
	templateTwo   = "2006-01-02T15:04:05+00:00"
	templateThree = "2006-01-02 15:04:05"
	templateFour  = "2006-01-02"
	loc, _        = time.LoadLocation("Asia/Shanghai") //设置时区
)

//  yyyy-mm-dd  hh:mm:ss 格式的时间转换为时间戳
func StringTimeToInt(str string) int {
	te, err := time.ParseInLocation(templateThree, str, loc)
	if err != nil {
		return 0
	}
	return int(te.Unix())
}

//  yyyy-mm-dd  hh:mm:ss 格式的时间转换为时间戳
func StringTimeToInt64(str string) int64 {
	te, err := time.ParseInLocation(templateThree, str, loc)
	if err != nil {
		return 0
	}
	return te.Unix()
}

//string转上海时间Time
func StringToTimeYMD(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ymd, _ := time.ParseInLocation(templateFour, str, loc)
	return ymd
}

//string转上海时间Time
func StringToTimeYMDHMS(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ymd, _ := time.ParseInLocation(templateThree, str, loc)
	return ymd
}

//时间戳转Time YYYY-MM-HH
func Int64ToYMD(t int64) time.Time {
	// go语言固定日期模版
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(t, 0).Format(templateFour)
	st, _ := time.Parse(templateFour, timeStr) //string转time
	return st
}

//时间戳转日期 YYYY-MM-HH MM:SS
func Int64ToDateTimeString(t int64) string {
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(t, 0).Format(templateThree)
	return timeStr
}

// 开始日期和结束日期相隔天数
func GetTimeArr(start, end string) int64 {
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(templateFour, start, loc)
	endUnix, _ := time.ParseInLocation(templateFour, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

//获得当前月的初始和结束日期
func GetMonthDay() (int64, int64) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Unix(), time.Unix(l, 0).Unix()
}

// 获取指定时间所在月的开始 结束时间
func GetMonthStartEnd(t time.Time) (int64, int64) {
	monthStartDay := t.AddDate(0, 0, -t.Day()+1)
	monthStartTime := time.Date(monthStartDay.Year(), monthStartDay.Month(), monthStartDay.Day(), 0, 0, 0, 0, t.Location())
	monthEndDay := monthStartTime.AddDate(0, 1, -1)
	monthEndTime := time.Date(monthEndDay.Year(), monthEndDay.Month(), monthEndDay.Day(), 23, 59, 59, 0, t.Location())
	return monthStartTime.Unix(), monthEndTime.Unix()
}