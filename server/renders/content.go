package render

type USER struct {
	ID        int
	User_name string
}
type Commentaire struct {
	user USER
	Date string
	Text string
}
type Post struct {
	ID_Post     int
	user        USER
	Title       string
	Category    string
	Description string
	Img         string
	comments    []Commentaire
}
