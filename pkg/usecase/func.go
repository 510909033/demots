package usecase

import "time"

// 获取今天开始时间戳
func GetTodayStartTs() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
}

// 获取今天结束时间戳
func GetTodayEndTs() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Unix()
}

// 获取本周的开始时间戳
func GetWeekStartTs() int64 {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(now.Year(), now.Month(), now.Day()-weekday+1, 0, 0, 0, 0, now.Location()).Unix()
}

// 获取本周的结束时间戳
func GetWeekEndTs() int64 {
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return time.Date(now.Year(), now.Month(), now.Day()-weekday+7, 23, 59, 59, 0, now.Location()).Unix()
}

// 获取本月的开始时间戳
func GetMonthStartTs() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
}

// 获取本月的结束时间戳
func GetMonthEndTs() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location()).Unix()
}
