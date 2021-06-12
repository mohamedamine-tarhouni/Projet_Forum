package render

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Render_Upload(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("UN")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("File Upload Endpoint Hit")
	//opening the database
	database, err := sql.Open("sqlite3", "./Forum_Final.db")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = database.Exec("PRAGMA journal_mode=WAL;")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `Image`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	filename := ""
	file, handler, err_file := r.FormFile("Image")
	Title := r.PostFormValue("Title")
	Description := r.PostFormValue("Description")

	if err_file == nil {
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err_temp := ioutil.TempFile("temp-images", "upload-*.png")
		if err_temp != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		// fmt.Fprintf(w, "Successfully Uploaded File\n")
		filename = tempFile.Name()
		filename = strings.ReplaceAll(filename, "\\", "/")
	}
	link := r.URL.Path + ".html"
	println(link)
	ID := Select_ID(database, c.Value)
	query_insert := `INSERT INTO Post (Title,Categorie,Description,ID_user,Image) VALUES (?, ?,?,?,?)`
	if (Title != "") || (Description != "") || (filename != "") {
		_, err_insert := database.Exec(query_insert, Title, r.URL.Path[1:], Description, ID, filename)
		http.Redirect(w, r, link, http.StatusFound)
		if err_insert != nil {
			println("erreur d'insertion")
			log.Fatalf(err.Error())
		}
	} else {
		link = "/Post_" + r.URL.Path[1:] + ".html"
		http.Redirect(w, r, link, http.StatusFound)
	}

}
