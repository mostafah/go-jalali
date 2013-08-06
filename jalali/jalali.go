package jalali

import (
	"fmt"
	"strings"
	"time"
)

var (
	gDaysInMonth                  = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	jDaysInMonth                  = [...]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	persianMonthNames             = [...]string{1: "فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور", "مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}
	persianMonthNamesPEnglish     = [...]string{1: "Farvardin", "Ordibehesht", "Khordad", "Tir", "Mordad", "Shahrivar", "Mehr", "Aban", "Azar", "Dey", "Bahman", "Esfand"}
	persianWeekdayNames           = [...]string{"شنبه", "یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه"}
	persianWeekdayNamesPEnglish   = [...]string{"Shanbeh", "YekShanbeh", "DoShanbeh", "SehShanbeh", "ChaharShanbeh", "PanjShanbeh", "JomE"}
	persianWeekdayNamesShort      = [...]string{"ش", "۱ش ", "۲ش", "۳ش", "۴ش", "۵ش", "ج "}
	persianMeridianIndicator      = [...]string{"قبل از ظهر", "بعد از ظهر"}
	persianMeridianIndicatorShort = [...]string{"ق.ظ", "ب.ظ"}
)

type Weekday int

const (
	Shanbeh Weekday = iota
	YekShanbeh
	DoShanbeh
	SehShanbeh
	ChaharShanbeh
	PanjShanbeh
	JomE
)

func (d Weekday) String() string {
	return persianWeekdayNamesPEnglish[d]
}

type Month int

const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

func (m Month) String() string {
	return persianMonthNamesPEnglish[m]
}

// Jtog converts year, month, day and optional hour, min, sec, nsec as a jalali date to time.Time (gregorian date)
func Jtog(year, month, day int, tt ...int) time.Time {
	ts := make([]int, 4)
	for i, t := range tt {
		ts[i] = t
	}
	gy, gm, gd := jalaliToGregorian(year, month, day)
	return time.Date(gy, time.Month(gm), gd, ts[0], ts[1], ts[2], ts[3], time.Local)
}

// Gtoj converts t (gregorian date) to jalali date and returns jyear, jmonth, jday
func Gtoj(t time.Time) (jyear, jmonth, jday int) {
	return gregorianToJalali(t.Year(), int(t.Month()), t.Day())
}

// IsLeap returns true if t is a leap year in jalali calendar.
// Info: http://www.tondering.dk/claus/cal/persian.php#leap
func IsLeap(t time.Time) bool {
	jyear, _, _ := Gtoj(t)
	switch jyear % 33 {
	case 6, 22, 17, 13, 9, 5, 1, 30:
		return true
	}
	return false
}

// Strftime formats time according to the directives in the given format string. Any text not listed as a directive will be
// passed through to the output string.
//
// Format meanings:
//
//		[%a]   The abbreviated weekday name (۳ش)
//		[%A]   The full weekday name (یکشنبه)
//		[%b]   The month name (اردیبهشت)
//		[%B]   The month name in pinglish (Ordibehesht)
//		[%d]   Day of the month (01..31)
//		[%e]   Day of the month (1..31)
//		[%j]   Day of the year (1..366)
//		[%m]   Month of the year, zero-padded (01..12)
//		[%_m]  Month of the year, blank-padded ( 1..12)
//		[%-m]  Month of the year, no-padded (1..12)
//		[%w]   Day of the week (Sunday is 0, 0..6)
//		[%x]   Preferred representation for the date alone, no time in format YY/M/D
//		[%y]   Year without a century (00..99)
//		[%Y]   Year with century
//		[%H]   Hour of the day, 24-hour clock (00..23)
//		[%I]   Hour of the day, 12-hour clock (01..12)
//		[%M]   Minute of the hour (00..59)
//		[%p]   Meridian indicator ("بعد از ظهر" or "قبل از ظهر")
//		[%S]   Second of the minute (00..60)
//		[%x]   1392/04/02
//		[%-x]  92/4/2
//		[%X]   Preferred representation for the time alone, no date
//		[%Z]   Time zone name
//		[%%]   Literal %'' character
//
// Example:
// 		jalali.Strftime("Printed on %Y/%m/%d", time.Now())   #=> "Printed on 1392/04/02"
func Strftime(format string, t time.Time) string {
	jyear, jmonth, jday := Gtoj(t)
	output := format
	yy := fmt.Sprintf("%d", jyear)[2:]
	hh := t.Hour()
	if t.Hour() >= 12 {
		hh = t.Hour() - 12
	}
	ampm := persianMeridianIndicatorShort[0]
	if t.Hour() >= 12 {
		ampm = persianMeridianIndicatorShort[1]
	}
	AMPM := persianMeridianIndicator[0]
	if t.Hour() >= 12 {
		AMPM = persianMeridianIndicator[1]
	}

	zone, _ := t.Zone()

	// X:=[("%02d" % @hour),("%02d" % @min),("%02d" % @sec)].join(":"))
	re := [][]string{
		{"%%", "SUBSTITUTION_MARKER"},
		{"%m", fmt.Sprintf("%02d", jmonth)},
		{"%_m", fmt.Sprintf("% 2d", jmonth)},
		{"%-m", fmt.Sprintf("%d", jmonth)},
		{"%a", persianWeekdayNamesShort[t.Weekday()]},
		{"%A", persianWeekdayNames[t.Weekday()]},
		{"%b", persianMonthNames[jmonth]},
		{"%B", persianMonthNamesPEnglish[jmonth]},
		{"%d", fmt.Sprintf("%02d", jday)},
		{"%e", fmt.Sprintf("%d", jday)},
		{"%Y", fmt.Sprintf("%d", jyear)},
		{"%y", yy},
		{"%j", fmt.Sprintf("%d", t.YearDay())},
		{"%H", fmt.Sprintf("%02d", t.Hour())},
		{"%I", fmt.Sprintf("%02d", hh)},
		{"%M", fmt.Sprintf("%02d", t.Minute())},
		{"%S", fmt.Sprintf("%02d", t.Second())},
		{"%p", ampm},
		{"%P", AMPM},
		{"%w", fmt.Sprintf("%d", t.Weekday())},
		{"%Z", zone},
		{"%X", fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())},
		{"%x", fmt.Sprintf("%d/%02d/%02d", jyear, t.Month(), t.Day())},
		{"%-x", fmt.Sprintf("%s/%d/%d", yy, t.Month(), t.Day())},
		{"SUBSTITUTION_MARKER", "%"},
	}
	for _, r := range re {
		output = strings.Replace(output, r[0], r[1], -1)
	}
	return output
}
