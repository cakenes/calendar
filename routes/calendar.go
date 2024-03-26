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
	currentYear, _, currentDay := now.Date()
	_, currentWeek := now.ISOWeek()

	year := controllers.GetYear(now, monthOffset)
	months := controllers.GetMonths(now, monthOffset)
	weekdays := controllers.GetWeekdays(weekdayOffset)
	weeks := controllers.GetWeeks(now, weekdayOffset, monthOffset)
	prevDays := controllers.GetDaysPrev(now, weekdayOffset, monthOffset)
	days := controllers.GetDaysCurrent(now, monthOffset)
	nextDays := controllers.GetDaysNext(len(prevDays), len(days))

	data := map[string]interface{}{
		"offset":      monthOffset,
		"currentDay":  currentDay,
		"currentWeek": currentWeek,
		"currentYear": currentYear,
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
