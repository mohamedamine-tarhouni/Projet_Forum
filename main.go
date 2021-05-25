package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tmpl := template.Must(template.ParseFiles("template/creation_compte.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Call to ParseForm makes form fields available.
		err := r.ParseForm()
		if err != nil {
			print("Error\n")
			// Handle error here via logging and then return
		}
		//  
		//les valeurs du formulaire
		Prenom := r.PostFormValue("name")
		Nom := r.PostFormValue("last_name")
		Addresse := r.PostFormValue("Address")
		MDP := r.PostFormValue("MDP")
		Mail := r.PostFormValue("Mail")

		//ouverture de la base (on la crée si elle n'existe pas)
		database, _ := sql.Open("sqlite3", "./Forum.db")

		//creation du table Utilisateur
		statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Utilisateur (ID       INTEGER PRIMARY KEY ASC AUTOINCREMENT,Nom      STRING  NOT NULL,PRENOM   STRING  NOT NULL,MAIL     STRING  UNIQUENOT NULL,PASSWORD STRING  NOT NULL,ADDRESSE STRING  NOT NULL);")
		//lancement de la requete précédente
		statement.Exec()

		//insertion des valeurs dans la base avec la requete INSERT INTO
		statement, _ = database.Prepare("INSERT INTO Utilisateur (Nom, PRENOM,MAIL,PASSWORD,ADDRESSE) VALUES (?, ?,?,?,?)")
		//on insère dans la base si les valeurs ne sont pas vide
		if Nom != "" && Prenom != "" && Mail != "" && MDP != "" && Addresse != "" {
			statement.Exec(Nom, Prenom, Mail, MDP, Addresse)
			// statement.Exec()
		}

		//execution de template
		tmpl.ExecuteTemplate(w, "creation", "") // we need to execute the template to recieve the data
	})

	//relation du go avec le dossier assets
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":900", nil)

}
