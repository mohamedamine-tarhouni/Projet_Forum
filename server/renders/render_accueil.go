package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate_accueil(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			Path:  "/",
		})
		http.Redirect(w, r, "/Accueil.html", http.StatusFound)
		return
	}
	if r.URL.Path == "/logout.html" {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			Path:  "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:   "UN",
			MaxAge: -1,
			Path:   "/",
		})
		http.Redirect(w, r, "/Accueil.html", http.StatusFound)
	}
	// println(c.Value)
	parsedTemplate, _ := template.ParseFiles("./template/Accueil.html")
	err_tmpl := parsedTemplate.Execute(w, c)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
