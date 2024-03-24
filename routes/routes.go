package routes

import (
	"html/template"
	"log"
	"main/middleware"
	"net/http"
)

var renderer *TemplateRenderer

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Setup() {
	renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	router := http.NewServeMux()

	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))

	CalendarRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := renderer.Render(w, "index", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"prev": "1",
			"next": "2",
		}

		if err := renderer.Render(w, "header", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	middleware := middleware.CombineMiddlewares(middleware.Logging)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware(router),
	}

	log.Fatal(server.ListenAndServe())

}
