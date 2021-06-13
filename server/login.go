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

//this function initiate the database when called by opening it and checking for errors
func initdatabase(database string) *sql.DB {

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	return db
}
func renderTemplate_creation(w http.ResponseWriter, r *http.Request) {

	//Errors is a struct that contains all of the insertion errors which will be used to verify the submitted values
	var Errors render.Errors
	parsedTemplate, _ := template.ParseFiles("./template/creation_compte.html")
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		print("Error\n")
		// Handle error here via logging
	}

	//all of the Form Values
	Prenom := r.PostFormValue("first_name")
	Nom := r.PostFormValue("last_name")
	User_name := r.PostFormValue("User_name")
	MDP := r.PostFormValue("MDP")
	CMDP := r.PostFormValue("CMDP")
	Mail := r.PostFormValue("Mail")
	Date := r.PostFormValue("date_naissance")
	Sexe := r.PostFormValue("genre")

	//this is the submit button
	button := r.PostFormValue("button_submit")

	//if the button has been clicked on once verif_insert will be true
	verif_insert := button != ""

	//initiating the database(it gets created if it dosent exist)
	database := initdatabase("./Forum_Final.db")
	//inserting the values in the database using INSERT INTO
	query_insert := `INSERT INTO Utilisateur (Nom, PRENOM,MAIL,PASSWORD,User_name,Birth_Date,genre) VALUES (?, ?,?,?,?,?,?)`

	//we start checking if we have to insert the values only when the user submited the form atleast once
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
	//we close the database when we are done
	defer database.Close()

	//we need to turn journal mode on to prevent the database from being locked
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
	err_tmpl := parsedTemplate.Execute(w, Errors)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}

}

//this function renders the page that contains the login form
func renderTemplate_login(w http.ResponseWriter, r *http.Request) {

	//we parse the template
	parsedTemplate, _ := template.ParseFiles("./template/login.html")
	database := initdatabase("./Forum_Final.db")
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		print("Error\n")
		// Handle error here via logging
	}

	//we get the Mail and the password submitted by the user as well as the password in the database(if the mail exists)
	MDP := r.PostFormValue("MDP")
	Mail := r.PostFormValue("Mail")

	//thats the password inside the database which we will get using the Mail(it will be "0" if the mail dosent exist)
	password := render.Select_password(database, Mail)

	//if the mail was found then we start verifying the password
	if password != "0" {

		//verif is the function that verifies a normal password with a Hashed password and returns a match(boolean)
		if cryptage.Verif(MDP, password) {
			//we get the username to put it in the cookie to be able to identify the connected user
			username := render.Select_Username(database, Mail)

			//we change the status to "1"("1-> logged-in "2"-> guest(annonymous))
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
			//after the user has been successfully logged-in he is redirected to the Home Page
			http.Redirect(w, r, "/Accueil.html", http.StatusFound)
			println("tout est bon")
		} else {
			println("faux mot de passe")
		}
	}
	//we close the database after we are done
	defer database.Close()
	err_tmpl := parsedTemplate.Execute(w, nil)
	if err_tmpl != nil {
		log.Println("Error executing template :", err_tmpl)
		return
	}
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
}

//we need to call Handlefunc to every link to render these pages
func Login() {
	render.Create_Data()
	http.HandleFunc("/", render.RenderTemplate_accueil)
	http.HandleFunc("/Accueil.html", render.RenderTemplate_accueil)
	http.HandleFunc("/creation_compte.html", renderTemplate_creation)
	http.HandleFunc("/login.html", renderTemplate_login)
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fs_img := http.FileServer(http.Dir("./temp-images/"))
	http.Handle("/temp-images/", http.StripPrefix("/temp-images/", fs_img))
}
