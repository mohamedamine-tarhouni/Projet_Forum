package render

import (
	"database/sql"
	"log"
	"strconv"
)

//this function returns the Username using his Mail
func Select_Username(db *sql.DB, address string) string {

	//we get the Username by his Mail
	query := "SELECT User_name FROM Utilisateur WHERE MAIL='" + address + "'"

	//we execute the previously written query
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		println("utilisateur n'existe pas")
	}
	var User string

	//we scan the columns for the username to return it
	for result.Next() {
		result.Scan(&User)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return User
	}
	return "0"
}

//this function returns the User_ID using his Username
func Select_ID(db *sql.DB, Username string) int {
	query := "SELECT ID_user FROM Utilisateur WHERE User_name='" + Username + "'"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		println("utilisateur n'existe pas")
	}
	var ID int
	for result.Next() {
		result.Scan(&ID)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return ID
	}
	return -1
}

//this function gets all the Posts related to a specific user and puts it in an arrays which is the result of this function
func Select_User_Posts(db *sql.DB, ID_user int) []Post {

	//we select the post by the user_ID
	query := "SELECT * FROM Post WHERE ID_user=?"
	result, _ := db.Query(query, ID_user)
	var posts []Post
	var ID int
	var Title string
	var Category string
	var Description string
	var Image string
	var ID_u int

	//we scan all the found data to put it inside the array which will be returned
	for result.Next() {
		var Post Post
		result.Scan(&ID, &Title, &Category, &Description, &ID_u, &Image)
		Post.ID_Post = ID
		Post.Title = Title
		Post.Description = Description
		Post.Category = Category
		Post.Comments = Select_comment(db, ID)
		Post.Img = Image
		// println(len(Post.Comments))
		Post.User = select_user(db, ID_user)
		posts = append(posts, Post)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	}
	return posts
}

//this function Selects the posts by Category
func Select_Posts(db *sql.DB, cat string) []Post {
	// query := "SELECT * FROM Post WHERE Categorie='informatique'"
	query := "SELECT * FROM Post WHERE Categorie='" + cat + "'"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		println("Post n'existe pas")
	}
	var posts []Post
	var ID int
	var Title string
	var Category string
	var Description string
	var Image string
	var ID_user int
	// i := 0
	for result.Next() {
		var Post Post
		result.Scan(&ID, &Title, &Category, &Description, &ID_user, &Image)
		Post.ID_Post = ID
		Post.Title = Title
		Post.Description = Description
		Post.Category = Category
		Post.Comments = Select_comment(db, ID)
		Post.Img = Image
		// println(len(Post.Comments))
		Post.User = select_user(db, ID_user)
		posts = append(posts, Post)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	}
	return posts
}

//this function returns the info about a user using his ID
func select_user(db *sql.DB, ID int) USER {
	ID_Str := strconv.Itoa(ID)

	query := "SELECT User_name FROM Utilisateur WHERE ID_user='" + ID_Str + "'"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		println("Post n'existe pas")
	}
	var Username string
	var User USER
	for result.Next() {
		result.Scan(&Username)
		User.ID = ID
		User.User_name = Username
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return User
	}
	return User
}

//this function returns a password using the Mail of the user or returns "0" if the mail dosen't exist
func Select_password(db *sql.DB, address string) string {
	query := "SELECT PASSWORD FROM Utilisateur WHERE MAIL='" + address + "'"
	result, err := db.Query(query)
	if err != nil {
		println("utilisateur n'existe pas")
	}
	var PASSWORD string
	for result.Next() {
		result.Scan(&PASSWORD)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return PASSWORD
	}
	return "0"
}

//this function returns an array containing all the comments related to a post using the Post_ID
func Select_comment(db *sql.DB, Post_id int) []Commentaire {
	var comments []Commentaire
	var commentaire Commentaire
	query := "SELECT * FROM Commentaire WHERE ID_Post=?"
	result, err := db.Query(query, Post_id)
	if err != nil {
		log.Fatal(err)
		println("Commentaire n'existe pas")
	}
	var ID_com int
	var ID_Post int
	var Date string
	var Text string
	var ID_User int
	for result.Next() {
		result.Scan(&ID_com, &ID_Post, &ID_User, &Date, &Text)
		// println("dans le for")
		commentaire.ID_Com = ID_com
		// println("COMMENT ID: ", commentaire.ID_Com)
		commentaire.Date = Date
		commentaire.Text = Text
		commentaire.User = select_user(db, ID_User)
		comments = append(comments, commentaire)
	}
	return comments
}
