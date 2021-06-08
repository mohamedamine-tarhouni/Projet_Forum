package render

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var Posts []Post
var comment string

func Render_Posts(w http.ResponseWriter, r *http.Request) {
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}

	link := "./template" + r.URL.Path
	println(r.URL.Path)
	parsedTemplate, _ := template.ParseFiles(link)
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	comment = r.PostFormValue("comment")
	println(comment)
	if comment != "" {
		http.Redirect(w, r, "/comment.html", http.StatusFound)
	}
	Posts = Select_Posts(database, "informatique")
	// i := 0
	// for range Posts {
	// 	println("ID POST : ", Posts[i].ID_Post)
	// 	println("ID USER : ", Posts[i].User.ID)
	// 	println("USERNAME : ", Posts[i].User.User_name)
	// 	println("CAT : ", Posts[i].Category)
	// 	println("Desc : ", Posts[i].Description)
	// 	println("Title : ", Posts[i].Title)
	// 	comments := Posts[i].Comments
	// 	j := 0
	// 	for range comments {
	// 		println("COMMENT ID :", comments[j].ID_Com)
	// 		j++
	// 	}
	// 	i++
	// }
	var data Data
	data.Posts = Posts
	data.Category = "informatique"
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
func Render_commenting(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}
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
	text := r.PostFormValue("comment")
	ID := Select_ID(database, c.Value)
	// println(comment)
	i, err := strconv.Atoi(comment)
	println("ID comment : ", i)
	query_insert := "INSERT INTO Commentaire (ID_post,ID_user,Texte,Date) VALUES (?, ?,?,?)"
	if text != "" {
		comment = ""
		_, err_insert := database.Exec(query_insert, i, ID, text, "10/10/2020")
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}
	}

	parsedTemplate, _ := template.ParseFiles("./template/comment.html")
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
