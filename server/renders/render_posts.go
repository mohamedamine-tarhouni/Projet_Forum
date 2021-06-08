package render

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

var Posts []Post

func Render_Posts(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	link := "./template" + r.URL.Path
	parsedTemplate, _ := template.ParseFiles(link)
	Posts = Select_Posts(database, "informatique")
	i := 0
	for range Posts {
		println("ID POST : ", Posts[i].ID_Post)
		println("ID USER : ", Posts[i].User.ID)
		println("USERNAME : ", Posts[i].User.User_name)
		println("CAT : ", Posts[i].Category)
		println("Desc : ", Posts[i].Description)
		println("Title : ", Posts[i].Title)
		i++
	}
	var data Data
	data.Posts = Posts
	err_tmpl := parsedTemplate.Execute(w, data)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
func Render_posting(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}
	// var Posting Post
	link := "./template" + r.URL.Path
	println(link)
	parsedTemplate, _ := template.ParseFiles(link)
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	Title := r.PostFormValue("Title")
	Description := r.PostFormValue("Description")
	ID := Select_ID(database, c.Value)
	println(Title)
	println(Description)
	println(ID)
	query_insert := `INSERT INTO Post (Title,Categorie,Description,ID_user) VALUES (?, ?,?,?)`
	if (Title != "") && (Description != "") {

		_, err_insert := database.Exec(query_insert, Title, "informatique", Description, ID)
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}
	}
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}
