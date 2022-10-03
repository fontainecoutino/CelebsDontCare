package templates

import (
	"html/template"
	"net/http"

	"github.com/fontainecoutino/CelebsDontCare/cors"
)

var tmpl *template.Template

// SetupRoutes
func SetupRoutes() {
	// be able to access assets folder
	/*
		fs := http.FileServer(http.Dir("assets"))
		http.Handle("/assets/", http.StripPrefix("/assets", fs))
	*/
	bootstrap := http.FileServer(http.Dir("bootstrap/assets"))
	http.Handle("/bootstrap/assets/", http.StripPrefix("/bootstrap/assets", bootstrap))

	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", assets))

	// home
	homeHandler := http.HandlerFunc(handleHome)
	http.Handle("/", cors.HtmlMiddleware(homeHandler))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
