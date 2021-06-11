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

func initdatabase(database string) *sql.DB {

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	return db
}
func renderTemplate_creation(w http.ResponseWriter, r *http.Request) {
	var Errors render.Errors
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
	CMDP := r.PostFormValue("CMDP")
	Mail := r.PostFormValue("Mail")
	Date := r.PostFormValue("date_naissance")
	Sexe := r.PostFormValue("genre")
	button := r.PostFormValue("button_submit")
	verif_insert := button != ""

	//ouverture de la base (on la crée si elle n'existe pas)
	database := initdatabase("./Forum_Final.db")
	//insertion des valeurs dans la base avec la requete INSERT INTO
	query_insert := `INSERT INTO Utilisateur (Nom, PRENOM,MAIL,PASSWORD,User_name,Birth_Date,genre) VALUES (?, ?,?,?,?,?,?)`

	// //on insère dans la base si les valeurs ne sont pas vide
	if verif_insert {
		Errors = render.Verif(Prenom, Nom, Mail, MDP, CMDP, User_name)
		if Errors.Err_name == "1" && Errors.Err_surname == "1" && Errors.Err_Email == "1" && Errors.Err_password == "1" && Errors.Err_User_name == "1" && Errors.Err_Cpassword == "1" {
			MDP_Hash, _ := cryptage.HashPassword(MDP)
			_, err_insert := database.Exec(query_insert, Nom, Prenom, Mail, MDP_Hash, User_name, Date, Sexe)
			if err_insert != nil {
				username_verif := render.Select_ID(database, User_name)
				if username_verif != -1 {
					Errors.Err_User_name = "3"
				} else {

					Errors.Err_Email = "3"
				}
			} else {
				http.Redirect(w, r, "/login.html", http.StatusFound)
			}
		}
	} else {
		Errors.Err_name = "1"
		Errors.Err_surname = "1"
		Errors.Err_User_name = "1"
		Errors.Err_Email = "1"
		Errors.Err_password = "1"
		Errors.Err_Cpassword = "1"
	}

	defer database.Close()
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
	// defer statement.Close()
	err_tmpl := parsedTemplate.Execute(w, Errors)
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
	password := render.Select_password(database, Mail)

	// println(password)
	if password != "0" {
		if cryptage.Verif(MDP, password) {
			username := render.Select_Username(database, Mail)
			http.SetCookie(w, &http.Cookie{
				Name:  "logged-in",
				Value: "1",
				Path:  "/",
			})
			http.SetCookie(w, &http.Cookie{
				Name:  "UN",
				Value: username,
				Path:  "/",
			})
			http.Redirect(w, r, "/Accueil.html", http.StatusFound)
			println("tout est bon")
		} else {
			println("faux mot de passe")
		}
	}
	defer database.Close()
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
}

func Login() {
	render.Create_Data()
	http.HandleFunc("/", render.RenderTemplate_accueil)
	http.HandleFunc("/Accueil.html", render.RenderTemplate_accueil)
	http.HandleFunc("/creation_compte.html", renderTemplate_creation)
	http.HandleFunc("/login.html", renderTemplate_login)
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
}
