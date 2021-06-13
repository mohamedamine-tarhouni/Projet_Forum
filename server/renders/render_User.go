package render

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func Render_User_Posts(w http.ResponseWriter, r *http.Request) {
	var i int
	cookie_user, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	link := "./template" + "/user_posts.html"
	parsedTemplate, _ := template.ParseFiles(link)
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	deletion := r.PostFormValue("delete")
	if deletion != "" {
		i, _ = strconv.Atoi(deletion)
	}
	database = Delete_Post_BY_ID(database, i)
	ID_user := Select_ID(database, cookie_user.Value)
	Posts = Select_User_Posts(database, ID_user)
	var data Data
	data.Posts = Posts
	err_tmpl := parsedTemplate.Execute(w, data)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
