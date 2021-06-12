package main

import (
	"log"
	"net/http"

	connexion "./server"

	_ "github.com/mattn/go-sqlite3"
)

const (
	Host = "localhost"
	Port = "2100"
)

func main() {
	connexion.Profil()
	connexion.Forum()
	connexion.Login()
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
