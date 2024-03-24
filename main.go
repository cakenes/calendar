package main

import (
	"database/sql"
	// "main/models"
	"main/routes"
	"net/http"
	"text/template"
)

var DB *sql.DB
var renderer *TemplateRenderer

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// models.Connect()
	// models.Setup()

	// defer models.DB.Close()

	routes.Setup()
	//
	// router := http.NewServeMux()
	//
	// year, month, day := time.Now().Date()
	// _, week := time.Now().ISOWeek()
	//
	// offset := 1
	//
	// weekdays := controllers.Weekdays(offset)
	// weeks := controllers.WeekNumbers(year, int(month), offset)
	// prevDays := controllers.DaysFromPreviousMonth(year, int(month), offset)
	// days := controllers.DaysFromCurrentMonth(year, int(month))
	// nextDays := controllers.DaysFromNextMonth(len(prevDays), len(days))
	//
	// router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	//
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if err := renderer.Render(w, "index", nil); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	//
	// router.HandleFunc("/calendar", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	//
	// 	if r.Method == http.MethodGet {
	// 		data := map[string]interface{}{
	// 			"day":      day,
	// 			"year":     year,
	// 			"weekdays": weekdays,
	// 			"week":     week,
	// 			"weeks":    weeks,
	// 			"prevDays": prevDays,
	// 			"days":     days,
	// 			"nextDays": nextDays,
	// 		}
	//
	// 		if err := renderer.Render(w, "calendar", data); err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		}
	// 	}
	// })
	//
	// router.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
	// 	data := map[string]interface{}{
	// 		prev:    int(month) - 1,
	// 		"month": int(month),
	// 		"year":  year,
	// 	}
	//
	// 	if err := renderer.Render(w, "header", data); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	//
	// middleware := middleware.CombineMiddlewares(middleware.Logging)
	//
	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: middleware(router),
	// }
	//
	// log.Fatal(server.ListenAndServe())

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
