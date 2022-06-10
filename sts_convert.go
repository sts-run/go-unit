package go_unit

import "strconv"

//2 标准库转换不够简洁，如string转int  有两个参数，使得代码不够简洁，这里一个封装，但必须确保数据是可预知的不出错误的值

//string转int
func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return value
}

//string转int32
func StringToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	value, _ := strconv.Atoi(s)
	return int32(value)
}

//string转int64
func StringToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}

//string转float64
func StringToFloat64(s string) float64 {
	value, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0.0
	}
	return value
}

//float64转string
func Float64ToStr(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}
