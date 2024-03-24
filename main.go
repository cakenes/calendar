package main

import (
	"database/sql"
	"log"
	"main/middleware"
	"main/models"
	"math"
	"net/http"
	"text/template"
	"time"
)

var DB *sql.DB
var renderer *TemplateRenderer

type TemplateRenderer struct {
	templates *template.Template
}

func Weekdays(startDay int) []string {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	adjustedDays := append(days[startDay:], days[:startDay]...)
	return adjustedDays
}

func WeekNumbers(year, month, startDay int) map[int]int {
	weekNumbers := make(map[int]int)
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

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

func DaysFromPreviousMonth(year, month, startDay int) []int {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	startDayOffset := (7 - startDay + int(firstDay.Weekday())) % 7
	prevMonthLastDay := firstDay.AddDate(0, 0, -1)
	prevMonthNumDays := prevMonthLastDay.Day()

	daysFromPrevMonth := make([]int, startDayOffset)
	for i := 0; i < startDayOffset; i++ {
		daysFromPrevMonth[i] = prevMonthNumDays - startDayOffset + i + 1
	}
	return daysFromPrevMonth
}

func DaysFromCurrentMonth(year, month int) []int {
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

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	models.Connect()
	models.Setup()

	defer models.DB.Close()

	renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.gohtml")),
	}

	router := http.NewServeMux()

	year, month, day := time.Now().Date()
	_, week := time.Now().ISOWeek()

	offset := 1

	weekdays := Weekdays(offset)
	weeks := WeekNumbers(year, int(month), offset)
	prevDays := DaysFromPreviousMonth(year, int(month), offset)
	days := DaysFromCurrentMonth(year, int(month))
	nextDays := DaysFromNextMonth(len(prevDays), len(days))

	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := renderer.Render(w, "index", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/calendar", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"day":      day,
			"month":    month,
			"year":     year,
			"weekdays": weekdays,
			"week":     week,
			"weeks":    weeks,
			"prevDays": prevDays,
			"days":     days,
			"nextDays": nextDays,
		}

		if err := renderer.Render(w, "calendar", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		if err := renderer.Render(w, "header", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	middleware := middleware.CombineMiddlewares(middleware.Logging)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware(router),
	}

	log.Fatal(server.ListenAndServe())

	// r.POST("/register", controllers.CreateUser)
	// r.POST("/login", controllers.LoginUser)
	// r.GET("/validate", controllers.CheckValidation)
	// r.Static("/assets", "./assets")
	// r.LoadHTMLGlob("templates/*")

	// r.Static("/css", "templates/css")

	// r.GET("/login", func(c *gin.Context) {
	// 	t := template.Must(template.ParseFS(templateFS,
	// 		"templates/login.html",
	// 	))
	// 	t.Execute(c.Writer, gin.H{"ts": timeNow()})
	// })

	// r.POST("/login", func(c *gin.Context) {
	// 	username := c.PostForm("username")
	// 	password := c.PostForm("password")
	//
	// 	if username == "admin" && password == "admin" {
	// 		println("Redirecting to /calendar")
	// 		c.SetCookie("sessionId", "0293845fn092m34im", 3600, "/", "localhost", false, true)
	// 		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// 	}
	// })

	// r.GET("/calendar", func(c *gin.Context) {
	// 	_, err := c.Cookie("sessionId")
	// 	if err != nil {
	// 		c.Redirect(302, "/login")
	// 	}
	//
	// 	t := template.Must(template.ParseFS(templateFS,
	// 		"templates/calendar.html",
	// 		"templates/calendar_title.html",
	// 	))
	// 	t.Execute(c.Writer, gin.H{"ts": timeNow()})
	// })

	// r.GET("/", func(c *gin.Context) {
	// 	cookie, err := c.Cookie("sessionId")
	// 	if err != nil {
	// 		c.SetCookie("sessionId", "0293845fn092m34im", 3600, "/", "localhost", false, true)
	// 	} else {
	// 		fmt.Println("Current cookie: ", cookie)
	// 	}

	// 	fmt.Println(c.Request.Header.Get("Authorization"))

	// 	t := template.Must(template.ParseFS(templateFS,
	// 		"templates/index.html",
	// 		"templates/login.html",
	// 		"templates/htmx.html",
	// 		"templates/calendar.html",
	// 		"templates/calendar_title.html"))
	// 	t.Execute(c.Writer, gin.H{"ts": timeNow()})
	// })

	// r.GET("/logout", func(c *gin.Context) {
	// 	c.SetCookie("sessionId", "", -1, "/", "localhost", false, true)
	// })

	// r.GET("/time", func(c *gin.Context) {
	// 	t := template.Must(template.ParseFS(templateFS,
	// 		"templates/htmx_time.html"))
	// 	t.Execute(c.Writer, gin.H{"ts": timeNow()})
	// })
}
