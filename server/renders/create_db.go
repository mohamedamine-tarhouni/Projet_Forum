package render

import (
	"database/sql"
	"log"
)

func Create_Data() {
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	query_user := `CREATE TABLE IF NOT EXISTS Utilisateur (
		ID_user    INTEGER            PRIMARY KEY ASC AUTOINCREMENT,
		Nom        STRING             NOT NULL,
		PRENOM     STRING             NOT NULL,
		MAIL       [STRING UNIQUENOT] NOT NULL
									  UNIQUE,
		PASSWORD   STRING             NOT NULL,
		User_name  STRING             NOT NULL
									  UNIQUE,
		Birth_Date DATE,
		genre      STRING
	);					`
	query_post := `CREATE TABLE IF NOT EXISTS Post (
		ID_post     INTEGER       PRIMARY KEY AUTOINCREMENT,
		Title       STRING (70)   NOT NULL,
		Categorie   STRING (30)   NOT NULL,
		Description STRING (2000) NOT NULL,
		ID_user                   REFERENCES Utilisateur (ID_user) ON DELETE CASCADE
																   ON UPDATE CASCADE   
	);					`
	query_react := `CREATE TABLE IF NOT EXISTS  Reaction (
		ID_react INTEGER  PRIMARY KEY AUTOINCREMENT,
		ID_user           REFERENCES Utilisateur (ID_user) ON DELETE CASCADE
														   ON UPDATE CASCADE,
		ID_Post           REFERENCES Post (ID_post) ON DELETE CASCADE
													ON UPDATE CASCADE,
		value    BOOLEAN  NOT NULL,
		date     DATETIME NOT NULL
	);					`
	query_comment := `CREATE TABLE IF NOT EXISTS Commentaire (
		ID_com  INTEGER       PRIMARY KEY AUTOINCREMENT,
		ID_post               REFERENCES Post (ID_post) ON DELETE CASCADE
														ON UPDATE CASCADE,
		ID_user               REFERENCES Utilisateur (ID_user) ON DELETE CASCADE
															   ON UPDATE CASCADE,
		Date    DATETIME,
		Texte   STRING (2000) NOT NULL
	);					`
	query_category := `CREATE TABLE IF NOT EXISTS  Categorie (
		ID_Cat  INTEGER PRIMARY KEY AUTOINCREMENT
						NOT NULL,
		Lib_Cat STRING  NOT NULL
	);					`
	//creation du table Utilisateur
	_, _ = database.Exec(query_user)
	//creation du table Reaction
	_, _ = database.Exec(query_react)
	//creation du table Post
	_, _ = database.Exec(query_post)
	//creation du table Commentaire
	_, _ = database.Exec(query_comment)
	//creation du table Categorie
	_, _ = database.Exec(query_category)
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
}
