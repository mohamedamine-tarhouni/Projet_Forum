package connexion

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	cryptage "./crypt"
	render "./renders"
	_ "github.com/mattn/go-sqlite3"
)

func select_password(db *sql.DB, address string) string {
	query := "SELECT PASSWORD FROM Utilisateur WHERE MAIL='" + address + "'"
	result, err := db.Query(query)
	if err != nil {
		println("utilisateur n'existe pas")
	}
	var PASSWORD string
	for result.Next() {
		result.Scan(&PASSWORD)
		defer db.Close()
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		return PASSWORD
	}
	return "Utilisateur n'existe pas dans la base"
}
func initdatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	return db
}
func renderTemplate_creation(w http.ResponseWriter, r *http.Request) {

	parsedTemplate, _ := template.ParseFiles("./template/creation_compte.html")
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	//les valeurs du formulaire
	Prenom := r.PostFormValue("first_name")
	Nom := r.PostFormValue("last_name")
	User_name := r.PostFormValue("User_name")
	MDP := r.PostFormValue("MDP")
	Mail := r.PostFormValue("Mail")
	Date := r.PostFormValue("date_naissance")
	Sexe := r.PostFormValue("genre")
	println(Date)
	println(Sexe)
	//ouverture de la base (on la crée si elle n'existe pas)
	database := initdatabase("./Forum_Final.db")
	//insertion des valeurs dans la base avec la requete INSERT INTO
	query_insert := `INSERT INTO Utilisateur (Nom, PRENOM,MAIL,PASSWORD,User_name,Birth_Date,genre) VALUES (?, ?,?,?,?,?,?)`
	if Nom != "" && Prenom != "" && Mail != "" && MDP != "" && User_name != "" {
		MDP_Hash, _ := cryptage.HashPassword(MDP)
		// //on insère dans la base si les valeurs ne sont pas vide
		_, err := database.Exec(query_insert, Nom, Prenom, Mail, MDP_Hash, User_name, Date, Sexe)
		if err != nil {
			println("erreur d'insertion")
			log.Fatal(err)
		}
	}
	defer database.Close()
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
	// defer statement.Close()
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}
func renderTemplate_login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./template/login.html")
	database := initdatabase("./Forum_Final.db")
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	MDP := r.PostFormValue("MDP")
	Mail := r.PostFormValue("Mail")
	password := select_password(database, Mail)
	defer database.Close()
	println(password)
	if cryptage.Verif(MDP, password) {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "1",
			Path:  "/",
		})
		http.Redirect(w, r, "/Accueil.html", http.StatusFound)
		println("tout est bon")
	} else {
		println("faux mot de passe")
	}
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
}

func Login() {
	// forum.forum()
	render.Create_Data()
	http.HandleFunc("/", render.RenderTemplate_accueil)
	http.HandleFunc("/Accueil.html", render.RenderTemplate_accueil)
	http.HandleFunc("/creation_compte.html", renderTemplate_creation)
	http.HandleFunc("/login.html", renderTemplate_login)
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
}
