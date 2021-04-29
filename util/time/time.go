/**
 *@Description
 *@ClassName time
 *@Date 2020/11/10 12:19 下午
 *@Author ckhero
 */

package time

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
	"youmi-micro-cluster/src/common/constant"
)

func GetBeginOfYesterday() time.Time{
	return now.BeginningOfDay().AddDate(0, 0, -1)
}

func GetBeginOfYesterdayStr() string {
	return GetBeginOfYesterday().Format(constant.DateTimeLayout)
}

func GetBeginOfDay() time.Time{
	return now.BeginningOfDay()
}

func GetBeginOfDayStr() string {
	return GetBeginOfDay().Format(constant.DateTimeLayout)
}

func IsEqualDayStr(first, second string) bool {
	f, _ := now.Parse(first)
	s, _ := now.Parse(second)
	return IsEqualDay(f, s)
}

func IsEqualDay(first, second time.Time) bool {
	return first.Year() == second.Year() && first.Month() == second.Month() && first.Day() == second.Day()
}

func GetDaysAfter(days int) time.Time{
	return now.New(time.Now()).AddDate(0, 0, days)
}

func GetDaysAfterStr(days int) string {
	return GetDaysAfter(days).Format(constant.DateTimeLayout)
}

func GetNowStr() string {
	return time.Now().Format(constant.DateTimeLayout)
}

func IsAfterStr(first, second string) bool {
	f, _ := now.Parse(first)
	s, _ := now.Parse(second)
	return IsAfter(f, s)
}

func IsAfter(first, second time.Time) bool {
	return now.New(first).After(second)
}

func GetDateAfter(days int) time.Time{
	return now.BeginningOfDay().AddDate(0, 0, days)
}

func GetDateAfterStr(days int) string {
	return GetDateAfter(days).Format(constant.DateTimeLayout)
}

func FormatDate(date string) string {
	t, _ := now.Parse(date)
	return now.New(t).Format(constant.DateLayout)
}

func FormatHour(date string) string {
	t, _ := now.Parse(date)
	return now.New(t).Format(constant.HourLayout)
}

func CurrDateTime() time.Time {
	return time.Now()
}

func CurrDateTimeStr() string {
	return CurrDateTime().Format(constant.DateTimeLayout)
}

func OneYearLater() time.Time{
	return GetDateAfter(365)
}

func Parse(str string) (time.Time, error) {
	return now.Parse(str)
}

func GetSecondsLeftToday() int {
	return int(now.BeginningOfDay().AddDate(0, 0, 1).Unix() - time.Now().Unix())
}

func EndOfDayStr(date string) string {
	t, _ := now.Parse(date)
	return now.New(t).EndOfDay().Format(constant.DateTimeLayout)
}

func BeginOfDayStr(date string) string {
	t, _ := now.Parse(date)
	return now.New(t).BeginningOfDay().Format(constant.DateTimeLayout)
}

func DiffTwo(first, second time.Time) int64 {
	return first.Unix() -  second.Unix()
}

func DiffTwoStr(first, second string) int64 {
	fmt.Println(first, second)
	f, _ := now.Parse(first)
	s, _ := now.Parse(second)
	return DiffTwo(f, s)
}

func SepToBeginAndEndStr(curr string) (string, string) {
	return BeginOfDayStr(curr), EndOfDayStr(curr)
}

