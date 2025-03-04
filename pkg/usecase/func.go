package usecase

import (
	"time"

	"github.com/jinzhu/now"
)

func init() {
	now.WeekStartDay = time.Monday
}

// 获取今天开始时间戳
func GetTodayStartTs() int64 {
	return now.BeginningOfDay().Unix()
}

// 获取今天结束时间戳
func GetTodayEndTs() int64 {
	return now.EndOfDay().Unix()
}

// 获取本周的开始时间戳
func GetWeekStartTs() int64 {
	return now.BeginningOfWeek().Unix()
}

// 获取本周的结束时间戳
func GetWeekEndTs() int64 {
	return now.EndOfWeek().Unix()
}

// 获取本月的开始时间戳
func GetMonthStartTs() int64 {
	return now.BeginningOfMonth().Unix()
}

// 获取本月的结束时间戳
func GetMonthEndTs() int64 {
	return now.EndOfMonth().Unix()
}
