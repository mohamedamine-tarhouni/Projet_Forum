package render

type USER struct {
	ID        int
	User_name string
}
type Commentaire struct {
	ID_Com int
	User   USER
	Date   string
	Text   string
}
type Post struct {
	ID_Post     int
	User        USER
	Title       string
	Category    string
	Description string
	Img         string
	Comments    []Commentaire
}
type Data struct {
	Posts    []Post
	Category string
	Status   string
}
type Errors struct {
	Err_name      string
	Err_surname   string
	Err_User_name string
	Err_Email     string
	Err_password  string
}
