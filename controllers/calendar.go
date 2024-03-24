package controllers

import (
	"math"
	"time"
)

func Weekdays(startDay int) []string {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	adjustedDays := append(days[startDay:], days[:startDay]...)
	return adjustedDays
}

func WeekNumbers(now time.Time, startDay, monthOffset int) map[int]int {
	now = now.AddDate(0, monthOffset, 0)
	year, month, _ := now.Date()
	weekNumbers := make(map[int]int)
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	offset := int(firstDay.Weekday()) - startDay
	if offset < 0 {
		offset += 7
	}
	adjustedFirstDay := firstDay.AddDate(0, 0, -offset)
	_, week := adjustedFirstDay.ISOWeek()

	for i := 2; i <= 7; i++ {
		weekNumbers[i] = week + i - 2
	}

	yearWeeks := weeksInYear(year) // Rolling over to the next year
	for i := 2; i <= 7; i++ {
		if weekNumbers[i] > yearWeeks {
			weekNumbers[i] = weekNumbers[i] - yearWeeks
		}
	}

	return weekNumbers
}

func weeksInYear(year int) int {
	lastDay := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	_, week := lastDay.ISOWeek()
	if week == 1 {
		return 52
	}
	return week
}

func DaysFromPreviousMonth(now time.Time, startDay, monthOffset int) []int {
	now = now.AddDate(0, monthOffset, 0)
	year, month, _ := now.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	startDayOffset := (7 - startDay + int(firstDay.Weekday())) % 7
	prevMonthLastDay := firstDay.AddDate(0, 0, -1)
	prevMonthNumDays := prevMonthLastDay.Day()
	daysFromPrevMonth := make([]int, startDayOffset)
	for i := 0; i < startDayOffset; i++ {
		daysFromPrevMonth[i] = prevMonthNumDays - startDayOffset + i + 1
	}
	return daysFromPrevMonth
}

func DaysFromCurrentMonth(now time.Time, monthOffset int) []int {
	year, month, _ := now.Date()
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	numDays := lastDay.Day()

	daysFromCurrentMonth := make([]int, numDays)
	for i := 0; i < numDays; i++ {
		daysFromCurrentMonth[i] = i + 1
	}
	return daysFromCurrentMonth
}

func DaysFromNextMonth(prev, current int) []int {
	remaining := int(math.Max(0, float64(42-prev-current)))
	daysFromNextMonth := make([]int, remaining)
	for i := range daysFromNextMonth {
		daysFromNextMonth[i] = i + 1
	}
	return daysFromNextMonth
}
