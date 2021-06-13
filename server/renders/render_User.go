package render

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//this function renders the page that contains all the user posts
func Render_User_Posts(w http.ResponseWriter, r *http.Request) {
	var i int

	//we need the cookie to identify the user
	cookie_user, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//opening the database
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}

	//we load the template
	link := "./template" + "/user_posts.html"
	parsedTemplate, _ := template.ParseFiles(link)
	err_forum := r.ParseForm()
	if err_forum != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}

	//this is the value of the delete button so the user can delete a post
	deletion := r.PostFormValue("delete")
	//when we click on the button the variable i will get the Post_Id the user wanted to delete
	if deletion != "" {
		i, _ = strconv.Atoi(deletion)
	}

	//we get all the data we need from the database and we display it
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
