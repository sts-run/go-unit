package go_unit

import (
	"fmt"
	"strings"
)

//	string切片 变成字符串 默认,号连接
func SliceIntToString(s []string, old, new string) string {
	if new == "" {
		new = ","
	}
	if old == "" {
		old = " "
	}
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1)
}

//	int64切片 变成字符串 默认,号连接
func SliceInt64ToString(s []int64, old, new string) string {
	if new == "" {
		new = ","
	}
	if old == "" {
		old = " "
	}
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), old, new, -1)
}

// IsEnglishText 检测字符串是否以字母开头
func IsEnglishText(str string) bool {
	u := str[0]
	if (u >= 65 && u <= 90) || (u >= 97 && u <= 122) {
		return true
	}
	return false
}