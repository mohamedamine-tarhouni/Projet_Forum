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

//this function render the page containing all the categories
func Render_Categories(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			Path:  "/",
		})
	}
	parsedTemplate, _ := template.ParseFiles("./template/Posts.html")
	err_tmpl := parsedTemplate.Execute(w, c)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}

//this function renders the page that contains all the posts related to a specific Category
func Render_Posts(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			Path:  "/",
		})
	}

	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}

	//this is the part where we get the categorie according to the link so we display the good page
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[1:]

	//the link to the template that contains the posts
	link := "./template" + "/categorie.html"
	parsedTemplate, _ := template.ParseFiles(link)
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}

	// this is from the button 'Post a new comment"
	comment = r.PostFormValue("comment")
	// println(comment)

	//if the user clicks the button he is redirected to the page where he will post a comment
	if comment != "" {
		comment_link := "/comment_" + Categorie + ".html"
		http.Redirect(w, r, comment_link, http.StatusFound)
	}

	//we select the posts according to the category
	Posts = Select_Posts(database, Categorie)

	//data is what we will need to display the proper posts in each category
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

//posting is the function that renders the page where the user is able to add a new post
func Render_posting(w http.ResponseWriter, r *http.Request) {

	//we need the category here aswell
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[6:] // (/->0 P->1 O->2 S-> 3 T ->4 _->5) thats why we take from the 6th position
	link := "./template" + "/Post_categorie.html"
	parsedTemplate, _ := template.ParseFiles(link)
	err_tmpl := parsedTemplate.Execute(w, Categorie)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}

//this function renders the page where the user is able to add a comment
func Render_commenting(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}

	//we need the category in the comments too
	category_link := r.URL.Path
	category_link = strings.ReplaceAll(category_link, ".html", "")
	Categorie := category_link[9:]
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

	//the comment written by the user
	text := r.PostFormValue("comment")

	//we need the user_ID to connect the comment to its proper writer
	ID := Select_ID(database, c.Value)
	i, err := strconv.Atoi(comment)

	//we insert the data in the comment to display it later
	query_insert := "INSERT INTO Commentaire (ID_post,ID_user,Texte,Date) VALUES (?, ?,?,?)"

	if text != "" {
		comment = ""

		//we execute the Sqlite3_request
		_, err_insert := database.Exec(query_insert, i, ID, text, "10/10/2020")
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}

		//after the user commented he gets redirected to whatever category they were in
		redirection_link := "/" + Categorie + ".html"
		http.Redirect(w, r, redirection_link, http.StatusFound)
	}
	//we load the page that contains the commenting textarea
	parsedTemplate, _ := template.ParseFiles("./template/comment.html")
	err_tmpl := parsedTemplate.Execute(w, Categorie)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
