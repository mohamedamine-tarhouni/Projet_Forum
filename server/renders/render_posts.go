package render

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var Posts []Post
var comment string

func Render_Categories(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./template/Posts.html")
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
func Render_Posts(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		log.Fatalf(err.Error())
	}
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[1:]
	println(Categorie)
	link := "./template" + "/categorie.html"
	parsedTemplate, _ := template.ParseFiles(link)
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	comment = r.PostFormValue("comment")
	// println(comment)
	if comment != "" {
		comment_link := "/comment_" + Categorie + ".html"
		http.Redirect(w, r, comment_link, http.StatusFound)
	}
	Posts = Select_Posts(database, Categorie)
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
	data.Category = Categorie
	data.Status = c.Value
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
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[6:]
	println(Categorie)
	// var Posting Post
	link := "./template" + "/Post_informatique.html"
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
		_, err_insert := database.Exec(query_insert, Title, Categorie, Description, ID)
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}
	}
	err_tmpl := parsedTemplate.Execute(w, Categorie)
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
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[9:]
	println(Categorie)
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
	query_insert := "INSERT INTO Commentaire (ID_post,ID_user,Texte,Date) VALUES (?, ?,?,?)"
	if text != "" {
		comment = ""
		_, err_insert := database.Exec(query_insert, i, ID, text, "10/10/2020")
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}
		redirection_link := "/" + Categorie + ".html"
		http.Redirect(w, r, redirection_link, http.StatusFound)
	}

	parsedTemplate, _ := template.ParseFiles("./template/comment.html")
	err_tmpl := parsedTemplate.Execute(w, Categorie)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
