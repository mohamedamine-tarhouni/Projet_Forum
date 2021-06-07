package render

import (
	"database/sql"
	"log"
	"strconv"
)

func Select_Username(db *sql.DB, address string) string {
	query := "SELECT User_name FROM Utilisateur WHERE MAIL='" + address + "'"
	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		println("utilisateur n'existe pas")
	}
	var User string
	for result.Next() {
		result.Scan(&User)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return User
	}
	return "0"
}
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

func Select_Posts(db *sql.DB, cat string) []Post {
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
	var ID_user int
	var Approval bool
	for result.Next() {
		var Post Post
		result.Scan(&ID, &Title, &Category, &Description, &ID_user, &Approval)
		Post.ID_Post = ID
		Post.Title = Title
		Post.Description = Description
		Post.user = select_user(db, ID_user)
		posts = append(posts, Post)
		// defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return posts
	}
	return nil
}

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
