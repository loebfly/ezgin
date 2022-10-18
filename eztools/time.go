package eztools

import (
	"database/sql/driver"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

// Time 是一个时间类型
/*
	使用说明：
		1. Time 是一个时间类型，用于替代 time.Time
		2. 已适配 json、bson、sql 的序列化和反序列化
		3. 提供了一些常用的时间操作方法
*/
type Time time.Time

func (t Time) OriVal() time.Time {
	return time.Time(t)
}

func (t Time) AddSeconds(seconds int) Time {
	return Time(t.OriVal().Add(time.Duration(seconds) * time.Second))
}

func (t Time) AddMinutes(minutes int) Time {
	return Time(t.OriVal().Add(time.Duration(minutes) * time.Minute))
}

func (t Time) AddHours(hours int) Time {
	return Time(t.OriVal().Add(time.Duration(hours) * time.Hour))
}

func (t Time) AddDays(days int) Time {
	return Time(t.OriVal().AddDate(0, 0, days))
}

func (t Time) AddMonths(months int) Time {
	return Time(t.OriVal().AddDate(0, months, 0))
}

func (t Time) AddYears(years int) Time {
	return Time(t.OriVal().AddDate(years, 0, 0))
}

func (t Time) SubtractSeconds(seconds int) Time {
	return Time(t.OriVal().Add(time.Duration(-seconds) * time.Second))
}

func (t Time) SubtractMinutes(minutes int) Time {
	return Time(t.OriVal().Add(time.Duration(-minutes) * time.Minute))
}

func (t Time) SubtractHours(hours int) Time {
	return Time(t.OriVal().Add(time.Duration(-hours) * time.Hour))
}

func (t Time) SubtractDays(days int) Time {
	return Time(t.OriVal().AddDate(0, 0, -days))
}

func (t Time) SubtractMonths(months int) Time {
	return Time(t.OriVal().AddDate(0, -months, 0))
}

func (t Time) SubtractYears(years int) Time {
	return Time(t.OriVal().AddDate(-years, 0, 0))
}

func (t Time) Add(d time.Duration) Time {
	return Time(t.OriVal().Add(d))
}

func (t Time) Sub(d time.Duration) Time {
	return Time(t.OriVal().Add(-d))
}

func (t Time) AddDate(years, months, days int) Time {
	return Time(t.OriVal().AddDate(years, months, days))
}

func (t Time) SubDate(years, months, days int) Time {
	return Time(t.OriVal().AddDate(-years, -months, -days))
}

func (t Time) DiffSeconds(t2 Time) int {
	return int(t.OriVal().Sub(t2.OriVal()).Seconds())
}

func (t Time) DiffMinutes(t2 Time) int {
	return int(t.OriVal().Sub(t2.OriVal()).Minutes())
}

func (t Time) DiffHours(t2 Time) int {
	return int(t.OriVal().Sub(t2.OriVal()).Hours())
}

func (t Time) DiffDays(t2 Time) int {
	return int(t.OriVal().Sub(t2.OriVal()).Hours() / 24)
}

func (t Time) Diff(t2 Time) time.Duration {
	return t.OriVal().Sub(t2.OriVal())
}

// Format 以任意格式获取日期时间字符串
// MMMM - month - January
// MMM - month - Jan
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func (t Time) Format(format string) string {
	newFmt := strings.Replace(format, "MMMM", "January", -1)
	newFmt = strings.Replace(newFmt, "MMM", "Jan", -1)
	newFmt = strings.Replace(newFmt, "MM", "01", -1)
	newFmt = strings.Replace(newFmt, "M", "1", -1)
	newFmt = strings.Replace(newFmt, "dddd", "Monday", -1)
	newFmt = strings.Replace(newFmt, "ddd", "Mon", -1)
	newFmt = strings.Replace(newFmt, "dd", "02", -1)
	newFmt = strings.Replace(newFmt, "d", "2", -1)
	newFmt = strings.Replace(newFmt, "yyyy", "2006", -1)
	newFmt = strings.Replace(newFmt, "yy", "06", -1)
	newFmt = strings.Replace(newFmt, "hh", "15", -1)
	newFmt = strings.Replace(newFmt, "HH", "03", -1)
	newFmt = strings.Replace(newFmt, "H", "3", -1)
	newFmt = strings.Replace(newFmt, "mm", "04", -1)
	newFmt = strings.Replace(newFmt, "m", "4", -1)
	newFmt = strings.Replace(newFmt, "ss", "05", -1)
	newFmt = strings.Replace(newFmt, "s", "5", -1)
	newFmt = strings.Replace(newFmt, "tt", "PM", -1)
	newFmt = strings.Replace(newFmt, "ZZZ", "MST", -1)
	newFmt = strings.Replace(newFmt, "Z", "Z07:00", -1)
	return t.OriVal().Format(newFmt)
}

/*** Json、mysql、mongo 适配器接口实现 ***/

func (t Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).UnixMilli())
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	var ts int64
	ts, err = strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	theTime := time.UnixMilli(ts)
	*t = Time(theTime)
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Time) GetBSON() (interface{}, error) {
	return time.Time(t).UnixMilli(), nil
}

func (t *Time) SetBSON(raw bson.Raw) error {
	var val int64
	if err := raw.Unmarshal(&val); err != nil {
		return err
	}
	*t = Time(time.UnixMilli(val))
	return nil
}
