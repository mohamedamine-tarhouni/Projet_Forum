package connexion

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	cryptage "./crypt"
	_ "github.com/mattn/go-sqlite3"
)

func select_password(db *sql.DB, address string) string {
	query := "SELECT * FROM Utilisateur WHERE MAIL='" + address + "'"
	result, err := db.Query(query)
	if err != nil {
		println("utilisateur n'existe pas")
	}
	var PASSWORD string
	var MAIL string
	var Nom string
	var PRENOM string
	var ID_user int
	var ADDRESSE string
	var Date string
	for result.Next() {
		result.Scan(&ID_user, &Nom, &PRENOM, &MAIL, &PASSWORD, &ADDRESSE, &Date)
		println("password = ", PASSWORD)
		return PASSWORD
	}
	return "Utilisateur n'existe pas dans la base"
}
func initdatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
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
	println(Date)
	//ouverture de la base (on la crée si elle n'existe pas)
	// database, err := sql.Open("sqlite3", "./Forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	database := initdatabase("./Forum.db")

	//creation du table Utilisateur
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Utilisateur (    ID_user    INTEGER            PRIMARY KEY ASC AUTOINCREMENT,Nom        STRING             NOT NULL,PRENOM     STRING             NOT NULL,MAIL       [STRING UNIQUENOT] NOT NULL							  UNIQUE,PASSWORD   STRING             NOT NULL,User_name  STRING             NOT NULL							  UNIQUE,Birth_Date DATE);")
	// //  lancement de la requete précédente
	// statement.Exec()

	//insertion des valeurs dans la base avec la requete INSERT INTO
	statement, _ := database.Prepare("INSERT INTO Utilisateur (Nom, PRENOM,MAIL,PASSWORD,User_name,Birth_Date) VALUES (?, ?,?,?,?,?)")
	MDP_Hash, _ := cryptage.HashPassword(MDP)
	//on insère dans la base si les valeurs ne sont pas vide
	if Nom != "" && Prenom != "" && Mail != "" && MDP != "" && User_name != "" {
		statement.Exec(Nom, Prenom, Mail, MDP_Hash, User_name, Date)
	}
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}
func renderTemplate_login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./template/login.html")
	database := initdatabase("./Forum.db")
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		print("Error\n")
		// Handle error here via logging and then return
	}
	MDP := r.PostFormValue("MDP")
	Mail := r.PostFormValue("Mail")
	// println(Mail)
	// println(MDP)
	password := select_password(database, Mail)
	println(password)
	if cryptage.Verif(MDP, password) {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "1",
			Path:  "/",
		})
		// http.Redirect(w, r, "/connected.html", http.StatusFound)
		println("tout est bon")
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			Path:  "/",
		})
		println("faux mot de passe")
		// http.Redirect(w, r, "/login.html", http.StatusFound)
	}
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}

func renderTemplate_verif(w http.ResponseWriter, r *http.Request) {

	parsedTemplate, _ := template.ParseFiles("./template/connected.html")
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}

func renderTemplate_accueil(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("logged-in")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	println(c.Value)
	parsedTemplate, _ := template.ParseFiles("./template/Accueil.html")
	err_tmpl := parsedTemplate.Execute(w, c)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
}
func Login() {
	http.HandleFunc("/", renderTemplate_accueil)
	http.HandleFunc("/Accueil.html", renderTemplate_accueil)
	http.HandleFunc("/creation_compte.html", renderTemplate_creation)
	http.HandleFunc("/login.html", renderTemplate_login)
	http.HandleFunc("/connected.html", renderTemplate_verif)
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

}
