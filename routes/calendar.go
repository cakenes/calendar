package routes

import (
	"main/controllers"
	"net/http"
	"strconv"
	"time"
)

func CalendarRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/calendar", calendar)
}

func calendar(w http.ResponseWriter, r *http.Request) {
	var monthOffset int = 0
	var weekdayOffset int = 1

	if err := r.ParseForm(); err != nil {
		println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if offset := r.FormValue("offset"); offset != "" {
		if parsedOffset, err := strconv.Atoi(offset); err == nil {
			monthOffset = parsedOffset
		}
	}

	if weekday := r.FormValue("weekday"); weekday != "" {
		if parsedWeekday, err := strconv.Atoi(weekday); err == nil {
			weekdayOffset = parsedWeekday
		}
	}

	now := time.Now()
	_, _, currentDay := now.Date()
	_, currentWeek := now.ISOWeek()

	offsetTime := now.AddDate(0, monthOffset, 0)
	year, _, _ := offsetTime.Date()

	months := controllers.GetMonths(now, monthOffset)
	weekdays := controllers.Weekdays(weekdayOffset)
	weeks := controllers.WeekNumbers(now, weekdayOffset, monthOffset)
	prevDays := controllers.DaysFromPreviousMonth(now, weekdayOffset, monthOffset)
	days := controllers.DaysFromCurrentMonth(now, monthOffset)
	nextDays := controllers.DaysFromNextMonth(len(prevDays), len(days))

	data := map[string]interface{}{
		"offset":      monthOffset,
		"offsetPrev":  monthOffset - 1,
		"offsetNext":  monthOffset + 1,
		"currentDay":  currentDay,
		"currentWeek": currentWeek,
		"months":      months,
		"year":        year,
		"weekdays":    weekdays,
		"weeks":       weeks,
		"prevDays":    prevDays,
		"days":        days,
		"nextDays":    nextDays,
	}

	if err := renderer.Render(w, "calendar", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
