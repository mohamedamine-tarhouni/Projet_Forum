package render

import (
	"log"
	"net/http"
	"text/template"
)

//this function renders the Home page
func RenderTemplate_accueil(w http.ResponseWriter, r *http.Request) {

	//we load a cookie and if it dosent exist we create it and we reload the page
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

	//if the user choose log out we remove the cookie containing his session and his Status becomes 0(disconnected)
	//and then we reload the page
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

	//we parse the Html file of the page Home
	parsedTemplate, _ := template.ParseFiles("./template/Accueil.html")

	//when we execute the template and we send the cookie so we display the site according to the user status
	err_tmpl := parsedTemplate.Execute(w, c)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
